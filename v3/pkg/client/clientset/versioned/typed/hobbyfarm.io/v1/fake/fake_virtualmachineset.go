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

// fakeVirtualMachineSets implements VirtualMachineSetInterface
type fakeVirtualMachineSets struct {
	*gentype.FakeClientWithList[*v1.VirtualMachineSet, *v1.VirtualMachineSetList]
	Fake *FakeHobbyfarmV1
}

func newFakeVirtualMachineSets(fake *FakeHobbyfarmV1, namespace string) hobbyfarmiov1.VirtualMachineSetInterface {
	return &fakeVirtualMachineSets{
		gentype.NewFakeClientWithList[*v1.VirtualMachineSet, *v1.VirtualMachineSetList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("virtualmachinesets"),
			v1.SchemeGroupVersion.WithKind("VirtualMachineSet"),
			func() *v1.VirtualMachineSet { return &v1.VirtualMachineSet{} },
			func() *v1.VirtualMachineSetList { return &v1.VirtualMachineSetList{} },
			func(dst, src *v1.VirtualMachineSetList) { dst.ListMeta = src.ListMeta },
			func(list *v1.VirtualMachineSetList) []*v1.VirtualMachineSet {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.VirtualMachineSetList, items []*v1.VirtualMachineSet) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
