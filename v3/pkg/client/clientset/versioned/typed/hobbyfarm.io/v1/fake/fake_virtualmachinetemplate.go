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

// fakeVirtualMachineTemplates implements VirtualMachineTemplateInterface
type fakeVirtualMachineTemplates struct {
	*gentype.FakeClientWithList[*v1.VirtualMachineTemplate, *v1.VirtualMachineTemplateList]
	Fake *FakeHobbyfarmV1
}

func newFakeVirtualMachineTemplates(fake *FakeHobbyfarmV1, namespace string) hobbyfarmiov1.VirtualMachineTemplateInterface {
	return &fakeVirtualMachineTemplates{
		gentype.NewFakeClientWithList[*v1.VirtualMachineTemplate, *v1.VirtualMachineTemplateList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("virtualmachinetemplates"),
			v1.SchemeGroupVersion.WithKind("VirtualMachineTemplate"),
			func() *v1.VirtualMachineTemplate { return &v1.VirtualMachineTemplate{} },
			func() *v1.VirtualMachineTemplateList { return &v1.VirtualMachineTemplateList{} },
			func(dst, src *v1.VirtualMachineTemplateList) { dst.ListMeta = src.ListMeta },
			func(list *v1.VirtualMachineTemplateList) []*v1.VirtualMachineTemplate {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.VirtualMachineTemplateList, items []*v1.VirtualMachineTemplate) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
