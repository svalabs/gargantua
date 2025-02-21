/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	fmt "fmt"

	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	v2 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v2"
	terraformcontrollercattleiov1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/terraformcontroller.cattle.io/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=hobbyfarm.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("accesscodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().AccessCodes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("costs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Costs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("courses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Courses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("dynamicbindconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().DynamicBindConfigurations().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("environments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Environments().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("onetimeaccesscodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().OneTimeAccessCodes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("predefinedservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().PredefinedServices().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("progresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Progresses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("scenarios"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Scenarios().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("scheduledevents"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().ScheduledEvents().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("scopes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Scopes().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("sessions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Sessions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("settings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Settings().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().Users().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().VirtualMachines().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachineclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().VirtualMachineClaims().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachinesets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().VirtualMachineSets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachinetemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V1().VirtualMachineTemplates().Informer()}, nil

		// Group=hobbyfarm.io, Version=v2
	case v2.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hobbyfarm().V2().Users().Informer()}, nil

		// Group=terraformcontroller.cattle.io, Version=v1
	case terraformcontrollercattleiov1.SchemeGroupVersion.WithResource("executions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Terraformcontroller().V1().Executions().Informer()}, nil
	case terraformcontrollercattleiov1.SchemeGroupVersion.WithResource("modules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Terraformcontroller().V1().Modules().Informer()}, nil
	case terraformcontrollercattleiov1.SchemeGroupVersion.WithResource("states"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Terraformcontroller().V1().States().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
