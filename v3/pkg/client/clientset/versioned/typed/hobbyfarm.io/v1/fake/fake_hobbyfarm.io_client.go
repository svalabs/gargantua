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
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHobbyfarmV1 struct {
	*testing.Fake
}

func (c *FakeHobbyfarmV1) AccessCodes(namespace string) v1.AccessCodeInterface {
	return newFakeAccessCodes(c, namespace)
}

func (c *FakeHobbyfarmV1) Costs(namespace string) v1.CostInterface {
	return newFakeCosts(c, namespace)
}

func (c *FakeHobbyfarmV1) Courses(namespace string) v1.CourseInterface {
	return newFakeCourses(c, namespace)
}

func (c *FakeHobbyfarmV1) DynamicBindConfigurations(namespace string) v1.DynamicBindConfigurationInterface {
	return newFakeDynamicBindConfigurations(c, namespace)
}

func (c *FakeHobbyfarmV1) Environments(namespace string) v1.EnvironmentInterface {
	return newFakeEnvironments(c, namespace)
}

func (c *FakeHobbyfarmV1) OneTimeAccessCodes(namespace string) v1.OneTimeAccessCodeInterface {
	return newFakeOneTimeAccessCodes(c, namespace)
}

func (c *FakeHobbyfarmV1) PredefinedServices(namespace string) v1.PredefinedServiceInterface {
	return newFakePredefinedServices(c, namespace)
}

func (c *FakeHobbyfarmV1) Progresses(namespace string) v1.ProgressInterface {
	return newFakeProgresses(c, namespace)
}

func (c *FakeHobbyfarmV1) Quizes(namespace string) v1.QuizInterface {
	return newFakeQuizes(c, namespace)
}

func (c *FakeHobbyfarmV1) Scenarios(namespace string) v1.ScenarioInterface {
	return newFakeScenarios(c, namespace)
}

func (c *FakeHobbyfarmV1) ScheduledEvents(namespace string) v1.ScheduledEventInterface {
	return newFakeScheduledEvents(c, namespace)
}

func (c *FakeHobbyfarmV1) Scopes(namespace string) v1.ScopeInterface {
	return newFakeScopes(c, namespace)
}

func (c *FakeHobbyfarmV1) Sessions(namespace string) v1.SessionInterface {
	return newFakeSessions(c, namespace)
}

func (c *FakeHobbyfarmV1) Settings(namespace string) v1.SettingInterface {
	return newFakeSettings(c, namespace)
}

func (c *FakeHobbyfarmV1) Users(namespace string) v1.UserInterface {
	return newFakeUsers(c, namespace)
}

func (c *FakeHobbyfarmV1) VirtualMachines(namespace string) v1.VirtualMachineInterface {
	return newFakeVirtualMachines(c, namespace)
}

func (c *FakeHobbyfarmV1) VirtualMachineClaims(namespace string) v1.VirtualMachineClaimInterface {
	return newFakeVirtualMachineClaims(c, namespace)
}

func (c *FakeHobbyfarmV1) VirtualMachineSets(namespace string) v1.VirtualMachineSetInterface {
	return newFakeVirtualMachineSets(c, namespace)
}

func (c *FakeHobbyfarmV1) VirtualMachineTemplates(namespace string) v1.VirtualMachineTemplateInterface {
	return newFakeVirtualMachineTemplates(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHobbyfarmV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
