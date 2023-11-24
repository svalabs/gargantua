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
	"context"

	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScheduledEvents implements ScheduledEventInterface
type FakeScheduledEvents struct {
	Fake *FakeHobbyfarmV1
	ns   string
}

var scheduledeventsResource = schema.GroupVersionResource{Group: "hobbyfarm.io", Version: "v1", Resource: "scheduledevents"}

var scheduledeventsKind = schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1", Kind: "ScheduledEvent"}

// Get takes name of the scheduledEvent, and returns the corresponding scheduledEvent object, and an error if there is any.
func (c *FakeScheduledEvents) Get(ctx context.Context, name string, options v1.GetOptions) (result *hobbyfarmiov1.ScheduledEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scheduledeventsResource, c.ns, name), &hobbyfarmiov1.ScheduledEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.ScheduledEvent), err
}

// List takes label and field selectors, and returns the list of ScheduledEvents that match those selectors.
func (c *FakeScheduledEvents) List(ctx context.Context, opts v1.ListOptions) (result *hobbyfarmiov1.ScheduledEventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scheduledeventsResource, scheduledeventsKind, c.ns, opts), &hobbyfarmiov1.ScheduledEventList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hobbyfarmiov1.ScheduledEventList{ListMeta: obj.(*hobbyfarmiov1.ScheduledEventList).ListMeta}
	for _, item := range obj.(*hobbyfarmiov1.ScheduledEventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scheduledEvents.
func (c *FakeScheduledEvents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scheduledeventsResource, c.ns, opts))

}

// Create takes the representation of a scheduledEvent and creates it.  Returns the server's representation of the scheduledEvent, and an error, if there is any.
func (c *FakeScheduledEvents) Create(ctx context.Context, scheduledEvent *hobbyfarmiov1.ScheduledEvent, opts v1.CreateOptions) (result *hobbyfarmiov1.ScheduledEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scheduledeventsResource, c.ns, scheduledEvent), &hobbyfarmiov1.ScheduledEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.ScheduledEvent), err
}

// Update takes the representation of a scheduledEvent and updates it. Returns the server's representation of the scheduledEvent, and an error, if there is any.
func (c *FakeScheduledEvents) Update(ctx context.Context, scheduledEvent *hobbyfarmiov1.ScheduledEvent, opts v1.UpdateOptions) (result *hobbyfarmiov1.ScheduledEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scheduledeventsResource, c.ns, scheduledEvent), &hobbyfarmiov1.ScheduledEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.ScheduledEvent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScheduledEvents) UpdateStatus(ctx context.Context, scheduledEvent *hobbyfarmiov1.ScheduledEvent, opts v1.UpdateOptions) (*hobbyfarmiov1.ScheduledEvent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scheduledeventsResource, "status", c.ns, scheduledEvent), &hobbyfarmiov1.ScheduledEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.ScheduledEvent), err
}

// Delete takes name of the scheduledEvent and deletes it. Returns an error if one occurs.
func (c *FakeScheduledEvents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scheduledeventsResource, c.ns, name, opts), &hobbyfarmiov1.ScheduledEvent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScheduledEvents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scheduledeventsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &hobbyfarmiov1.ScheduledEventList{})
	return err
}

// Patch applies the patch and returns the patched scheduledEvent.
func (c *FakeScheduledEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hobbyfarmiov1.ScheduledEvent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scheduledeventsResource, c.ns, name, pt, data, subresources...), &hobbyfarmiov1.ScheduledEvent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.ScheduledEvent), err
}
