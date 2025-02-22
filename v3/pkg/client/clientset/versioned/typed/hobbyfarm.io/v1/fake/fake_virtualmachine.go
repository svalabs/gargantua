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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeVirtualMachines implements VirtualMachineInterface
type fakeVirtualMachines struct {
	*gentype.FakeClientWithList[*v1.VirtualMachine, *v1.VirtualMachineList]
	Fake *FakeHobbyfarmV1
}

func newFakeVirtualMachines(fake *FakeHobbyfarmV1, namespace string) hobbyfarmiov1.VirtualMachineInterface {
	return &fakeVirtualMachines{
		gentype.NewFakeClientWithList[*v1.VirtualMachine, *v1.VirtualMachineList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("virtualmachines"),
			v1.SchemeGroupVersion.WithKind("VirtualMachine"),
			func() *v1.VirtualMachine { return &v1.VirtualMachine{} },
			func() *v1.VirtualMachineList { return &v1.VirtualMachineList{} },
			func(dst, src *v1.VirtualMachineList) { dst.ListMeta = src.ListMeta },
			func(list *v1.VirtualMachineList) []*v1.VirtualMachine { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.VirtualMachineList, items []*v1.VirtualMachine) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
