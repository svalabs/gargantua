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

// fakePredefinedServices implements PredefinedServiceInterface
type fakePredefinedServices struct {
	*gentype.FakeClientWithList[*v1.PredefinedService, *v1.PredefinedServiceList]
	Fake *FakeHobbyfarmV1
}

func newFakePredefinedServices(fake *FakeHobbyfarmV1, namespace string) hobbyfarmiov1.PredefinedServiceInterface {
	return &fakePredefinedServices{
		gentype.NewFakeClientWithList[*v1.PredefinedService, *v1.PredefinedServiceList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("predefinedservices"),
			v1.SchemeGroupVersion.WithKind("PredefinedService"),
			func() *v1.PredefinedService { return &v1.PredefinedService{} },
			func() *v1.PredefinedServiceList { return &v1.PredefinedServiceList{} },
			func(dst, src *v1.PredefinedServiceList) { dst.ListMeta = src.ListMeta },
			func(list *v1.PredefinedServiceList) []*v1.PredefinedService {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.PredefinedServiceList, items []*v1.PredefinedService) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
