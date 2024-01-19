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
	v2 "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/typed/hobbyfarm.io/v2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHobbyfarmV2 struct {
	*testing.Fake
}

func (c *FakeHobbyfarmV2) Scenarios(namespace string) v2.ScenarioInterface {
	return &FakeScenarios{c, namespace}
}

func (c *FakeHobbyfarmV2) ScheduledEvents(namespace string) v2.ScheduledEventInterface {
	return &FakeScheduledEvents{c, namespace}
}

func (c *FakeHobbyfarmV2) Users(namespace string) v2.UserInterface {
	return &FakeUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHobbyfarmV2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
