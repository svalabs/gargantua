package tfpcontroller

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	tfv1 "github.com/hobbyfarm/gargantua/pkg/apis/terraformcontroller.cattle.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	hfListers "github.com/hobbyfarm/gargantua/pkg/client/listers/hobbyfarm.io/v1"
	tfListers "github.com/hobbyfarm/gargantua/pkg/client/listers/terraformcontroller.cattle.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/util"
	k8sv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/retry"
	"k8s.io/client-go/util/workqueue"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type TerraformProvisionerController struct {
	hfClientSet *hfClientset.Clientset
	vmWorkqueue workqueue.RateLimitingInterface

	k8sClientset *k8s.Clientset

	tfClientset *hfClientset.Clientset

	vmLister  hfListers.VirtualMachineLister
	envLister hfListers.EnvironmentLister
	vmtLister hfListers.VirtualMachineTemplateLister

	tfsLister tfListers.StateLister
	tfeLister tfListers.ExecutionLister

	vmSynced  cache.InformerSynced
	vmtSynced  cache.InformerSynced
	tfsSynced  cache.InformerSynced
	tfeSynced  cache.InformerSynced
	envSynced  cache.InformerSynced
}

const (
	provisionNS = "hobbyfarm"
)

func NewTerraformProvisionerController(k8sClientSet *k8s.Clientset, hfClientSet *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) (*TerraformProvisionerController, error) {
	tfpController := TerraformProvisionerController{}
	tfpController.hfClientSet = hfClientSet

	tfpController.tfClientset = hfClientSet
	tfpController.k8sClientset = k8sClientSet

	tfpController.vmSynced = hfInformerFactory.Hobbyfarm().V1().VirtualMachines().Informer().HasSynced
	tfpController.vmtSynced = hfInformerFactory.Hobbyfarm().V1().VirtualMachineTemplates().Informer().HasSynced
	tfpController.envSynced = hfInformerFactory.Hobbyfarm().V1().Environments().Informer().HasSynced
	tfpController.tfsSynced = hfInformerFactory.Terraformcontroller().V1().States().Informer().HasSynced
	tfpController.tfeSynced = hfInformerFactory.Terraformcontroller().V1().Executions().Informer().HasSynced
	tfpController.vmWorkqueue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "VM")

	tfpController.vmLister = hfInformerFactory.Hobbyfarm().V1().VirtualMachines().Lister()
	tfpController.envLister = hfInformerFactory.Hobbyfarm().V1().Environments().Lister()
	tfpController.vmtLister = hfInformerFactory.Hobbyfarm().V1().VirtualMachineTemplates().Lister()

	tfpController.tfsLister = hfInformerFactory.Terraformcontroller().V1().States().Lister()
	tfpController.tfeLister = hfInformerFactory.Terraformcontroller().V1().Executions().Lister()

	vmInformer := hfInformerFactory.Hobbyfarm().V1().VirtualMachines().Informer()

	vmInformer.AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		AddFunc: tfpController.enqueueVM,
		UpdateFunc: func(old, new interface{}) {
			tfpController.enqueueVM(new)
		},
		DeleteFunc: tfpController.enqueueVM,
	}, time.Second*30)

	return &tfpController, nil
}

func (t *TerraformProvisionerController) enqueueVM(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		//utilruntime.HandleError(err)
		return
	}
	glog.V(8).Infof("Enqueueing vm %s", key)
	t.vmWorkqueue.AddRateLimited(key)
}

func (t *TerraformProvisionerController) Run(stopCh <-chan struct{}) error {
	defer t.vmWorkqueue.ShutDown()

	glog.V(4).Infof("Starting environment controller")
	glog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, t.vmSynced, t.envSynced, t.tfsSynced, t.tfeSynced, t.vmtSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}
	glog.Info("Starting environment controller worker")
	go wait.Until(t.runTFPWorker, time.Second, stopCh)
	glog.Info("Started environment controller worker")
	//if ok := cache.WaitForCacheSync(stopCh, )
	<-stopCh
	return nil
}

func (t *TerraformProvisionerController) runTFPWorker() {
	for t.processNextVM() {

	}
}

func (t *TerraformProvisionerController) processNextVM() bool {
	obj, shutdown := t.vmWorkqueue.Get()

	glog.V(4).Infof("processing VM")

	if shutdown {
		return false
	}
	err := func() error {
		defer t.vmWorkqueue.Done(obj)
		glog.V(4).Infof("processing vm in env controller: %v", obj)
		_, objName, err := cache.SplitMetaNamespaceKey(obj.(string)) // this is actually not necessary because VM's are not namespaced yet...
		if err != nil {
			glog.Errorf("error while splitting meta namespace key %v", err)
			//e.vmWorkqueue.AddRateLimited(obj)
			return nil
		}

		vm, err := t.vmLister.Get(objName)
		if err != nil {
			glog.Errorf("error while retrieving virtual machine %s, likely deleted %v", objName, err)
			t.vmWorkqueue.Forget(obj)
			return nil
		}

		err = t.handleProvision(vm)

		if err != nil {
			t.vmWorkqueue.AddRateLimited(obj)
			glog.Error(err)
		}
		t.vmWorkqueue.Forget(obj)
		glog.V(4).Infof("vm processed by endpoint controller %v", objName)

		return nil

	}()

	if err != nil {
		return true
	}

	return true
}

func (t *TerraformProvisionerController) handleProvision(vm *hfv1.VirtualMachine) error {
	if vm.Spec.Provision {
		glog.V(5).Infof("vm spec was to provision %s", vm.Name)
		if vm.DeletionTimestamp != nil {
			glog.V(5).Infof("destroying virtual machine")
			if vm.Status.TFState == "" {
				// vm already deleted let's delete our finalizer
				t.removeFinalizer(vm)
			}

			stateDel := t.tfClientset.TerraformcontrollerV1().States(provisionNS).Delete(vm.Status.TFState, &metav1.DeleteOptions{})
			if stateDel != nil {
				t.removeFinalizer(vm)
			}
		}
		if vm.Status.Status == "readyforprovisioning" {
			vmt, err := t.vmtLister.Get(vm.Spec.VirtualMachineTemplateId)
			if err != nil {
				glog.Errorf("error getting vmt %v", err)
				return err
			}
			env, err := t.envLister.Get(vm.Status.EnvironmentId)
			if err != nil {
				glog.Errorf("error getting env %v", err)
				return err
			}
			// let's provision the vm
			pubKey, privKey, err := util.GenKeyPair()
			if err != nil {
				glog.Errorf("error generating keypair %v", err)
				return err
			}
			envSpecificConfigFromEnv := env.Spec.EnvironmentSpecifics
			envTemplateInfo, exists := env.Spec.TemplateMapping[vmt.Name]
			if !exists {
				glog.Errorf("error pulling environment template info %v", err)
				return fmt.Errorf("environment template info does not exist for this template %s", vmt.Name)
			}
			config := make(map[string]string)
			for k, v := range envSpecificConfigFromEnv {
				config[k] = v
			}

			config["name"] = vm.Name
			config["public_key"] = pubKey
			config["cpu"] = strconv.Itoa(vmt.Spec.Resources.CPU)
			config["memory"] = strconv.Itoa(vmt.Spec.Resources.Memory)
			config["disk"] = strconv.Itoa(vmt.Spec.Resources.Storage)
			image, exists := envTemplateInfo["image"]
			if !exists{
				glog.Errorf("image does not exist in env template")
				return fmt.Errorf("image did not exist")
			}
			config["image"] = image

			r := fmt.Sprintf("%08x", rand.Uint32())
			cm := &k8sv1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name: strings.Join([]string{vm.Name + "-cm", r}, "-"),
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion: "v1",
							Kind: "VirtualMachine",
							Name: vm.Name,
							UID: vm.UID,
						},
					},
				},
				Data: config,
			}

			cm, err = t.k8sClientset.CoreV1().ConfigMaps(provisionNS).Create(cm)

			if err != nil {
				glog.Errorf("error creating configmap %s: %v", cm.Name, err)
			}

			keypair := &k8sv1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name: strings.Join([]string{vm.Name + "-secret", r}, "-"),
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion: "v1",
							Kind: "VirtualMachine",
							Name: vm.Name,
							UID: vm.UID,
						},
					},
				},
				Data: map[string][]byte{
					"private_key": []byte(privKey),
					"public_key":  []byte(pubKey),
				},
			}

			keypair, err = t.k8sClientset.CoreV1().Secrets(provisionNS).Create(keypair)

			if err != nil {
				glog.Errorf("error creating secret %s: %v", keypair.Name, err)
			}

			moduleName, exists := envTemplateInfo["module"]
			if !exists {
				glog.Errorf("module name does not exist")
				return fmt.Errorf("module name does not exist")
			}

			executorImage, exists := envTemplateInfo["executor_image"]
			if !exists {
				glog.Errorf("executor image does not exist")
				return fmt.Errorf("executorimage does not exist")
			}

			tfs := &tfv1.State{
				ObjectMeta: metav1.ObjectMeta{
					Name: strings.Join([]string{vm.Name + "-tfs", r}, "-"),
				},
				Spec: tfv1.StateSpec{
					Variables: tfv1.Variables{
						ConfigNames: []string{cm.Name},
					},
					Image: executorImage,
					AutoConfirm:     true,
					DestroyOnDelete: true,
					ModuleName: moduleName,
				},
			}

			tfs, err = t.tfClientset.TerraformcontrollerV1().States(provisionNS).Create(tfs)

			if err != nil {
				glog.Errorf("error creating tfs %v", err)
			}

			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				toUpdate, err := t.vmLister.Get(vm.Name)
				if err != nil {
					if apierrors.IsNotFound(err) {
						return nil
					} else {
						glog.Errorf("unknown error encountered when setting terminating %v", err)
					}
				}
				toUpdate.Spec.KeyPair = keypair.Name
				toUpdate.Status.Status = "provisioned"
				toUpdate.Status.TFState = tfs.Name
				toUpdate.Finalizers = []string{"tfp.controllers.hobbyfarm.io"}

				toUpdate, updateErr := t.hfClientSet.HobbyfarmV1().VirtualMachines().Update(toUpdate)
				if err := t.verifyVM(toUpdate); err != nil {
					glog.Errorf("error while verifying machine!!! %s", toUpdate.Name)
				}
				return updateErr
			})

			if retryErr != nil {
				return retryErr
			}
			glog.V(5).Infof("provisioned vm %s", vm.Name)
			return nil

		} else if vm.Status.Status == "provisioned" {
			// let's check the status of our tf provision
			tfState, err := t.tfsLister.States(provisionNS).Get(vm.Status.TFState)
			if err != nil {
				if apierrors.IsNotFound(err) {
					return fmt.Errorf("execution not found")
				}
				return nil
			}

			tfExec, err := t.tfeLister.Executions(provisionNS).Get(tfState.Status.ExecutionName)
			if err != nil {
				//glog.Error(err)
				if apierrors.IsNotFound(err) {
					return fmt.Errorf("execution not found")
				}
				return nil
			}

			if tfExec.Status.Outputs == "" {
				return fmt.Errorf("execution output was empty")
			}
			var tfOutput map[string]map[string]string

			err = json.Unmarshal([]byte(tfExec.Status.Outputs), &tfOutput)
			if err != nil {
				glog.Error(err)
			}
			env, err := t.envLister.Get(vm.Status.EnvironmentId)
			if err != nil {
				glog.Error(err)
				return fmt.Errorf("error getting environment")
			}
			glog.V(5).Infof("private ip is: %s", tfOutput["private_ip"]["value"])

			retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
				toUpdate, err := t.vmLister.Get(vm.Name)
				if err != nil {
					if apierrors.IsNotFound(err) {
						return nil
					} else {
						glog.Errorf("unknown error encountered when setting terminating %v", err)
					}
				}

				toUpdate.Status.PrivateIP = tfOutput["private_ip"]["value"]
				if _, exists := tfOutput["public_ip"]; exists {
					toUpdate.Status.PublicIP = tfOutput["public_ip"]["value"]
				} else {
					toUpdate.Status.PublicIP = translatePrivToPub(env.Spec.IPTranslationMap,tfOutput["private_ip"]["value"])
				}
				toUpdate.Status.Hostname = tfOutput["hostname"]["value"]
				toUpdate.Status.Status = "running"

				toUpdate, updateErr := t.hfClientSet.HobbyfarmV1().VirtualMachines().Update(toUpdate)
				if err := t.verifyVM(toUpdate); err != nil {
					glog.Errorf("error while verifying machine!!! %s", toUpdate.Name)
				}
				return updateErr
			})

			if retryErr != nil {
				return retryErr
			}

		}
	} else {
		glog.V(5).Infof("vm %s was not a provisioned vm", vm.Name)
	}
	//glog.Infof("no provision action required here")
	return nil
}

func translatePrivToPub(translationMap map[string]string, priv string) string {
	splitIp := strings.Split(priv, ".")

	origPrefix := splitIp[0]+"."+splitIp[1]+"."+splitIp[2]

	translation, ok := translationMap[origPrefix]

	if ok {
		return translation + "." + splitIp[3]
	}
	return ""

}

type tfOutput struct {

}

func (t *TerraformProvisionerController) removeFinalizer(vm *hfv1.VirtualMachine) error {
	glog.V(5).Infof("removing finalizer for vm %s", vm.Name)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		toUpdate, err := t.vmLister.Get(vm.Name)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return nil
			} else {
				glog.Errorf("unknown error encountered when setting terminating %v", err)
			}
		}
		if reflect.DeepEqual(toUpdate.Finalizers, []string{}) {
			return nil
		}
		toUpdate.Finalizers = []string{}
		glog.V(5).Infof("removing vm finalizer for %s", vm.Name)
		toUpdate, updateErr := t.hfClientSet.HobbyfarmV1().VirtualMachines().Update(toUpdate)
		if err := t.verifyVM(toUpdate); err != nil {
			glog.Errorf("error while verifying machine!!! %s", toUpdate.Name)
		}
		return updateErr
	})

	if retryErr != nil {
		glog.Errorf("error while updating vm object while setting terminating %v", retryErr)
	}
	return retryErr
}

func (t *TerraformProvisionerController) verifyVM(vm *hfv1.VirtualMachine) error {
	var err error
	glog.V(5).Infof("Verifying vm %s", vm.Name)
	for i := 0; i < 25; i++ {
		var fromCache *hfv1.VirtualMachine
		fromCache, err = t.vmLister.Get(vm.Name)
		if err != nil {
			if apierrors.IsNotFound(err) {
				break
			}
			glog.Error(err)
			return err
		}
		if util.ResourceVersionAtLeast(fromCache.ResourceVersion, vm.ResourceVersion) {
			glog.V(5).Infof("resource version matched for %s", vm.Name)
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	glog.Errorf("resource version didn't match for in time%s", vm.Name)
	return nil

}