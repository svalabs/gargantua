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

package v1

import (
	"context"
	"time"

	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	scheme "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DynamicBindConfigurationsGetter has a method to return a DynamicBindConfigurationInterface.
// A group's client should implement this interface.
type DynamicBindConfigurationsGetter interface {
	DynamicBindConfigurations(namespace string) DynamicBindConfigurationInterface
}

// DynamicBindConfigurationInterface has methods to work with DynamicBindConfiguration resources.
type DynamicBindConfigurationInterface interface {
	Create(ctx context.Context, dynamicBindConfiguration *v1.DynamicBindConfiguration, opts metav1.CreateOptions) (*v1.DynamicBindConfiguration, error)
	Update(ctx context.Context, dynamicBindConfiguration *v1.DynamicBindConfiguration, opts metav1.UpdateOptions) (*v1.DynamicBindConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DynamicBindConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DynamicBindConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DynamicBindConfiguration, err error)
	DynamicBindConfigurationExpansion
}

// dynamicBindConfigurations implements DynamicBindConfigurationInterface
type dynamicBindConfigurations struct {
	client rest.Interface
	ns     string
}

// newDynamicBindConfigurations returns a DynamicBindConfigurations
func newDynamicBindConfigurations(c *HobbyfarmV1Client, namespace string) *dynamicBindConfigurations {
	return &dynamicBindConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dynamicBindConfiguration, and returns the corresponding dynamicBindConfiguration object, and an error if there is any.
func (c *dynamicBindConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DynamicBindConfigurations that match those selectors.
func (c *dynamicBindConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DynamicBindConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DynamicBindConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dynamicBindConfigurations.
func (c *dynamicBindConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dynamicBindConfiguration and creates it.  Returns the server's representation of the dynamicBindConfiguration, and an error, if there is any.
func (c *dynamicBindConfigurations) Create(ctx context.Context, dynamicBindConfiguration *v1.DynamicBindConfiguration, opts metav1.CreateOptions) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dynamicBindConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dynamicBindConfiguration and updates it. Returns the server's representation of the dynamicBindConfiguration, and an error, if there is any.
func (c *dynamicBindConfigurations) Update(ctx context.Context, dynamicBindConfiguration *v1.DynamicBindConfiguration, opts metav1.UpdateOptions) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		Name(dynamicBindConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dynamicBindConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dynamicBindConfiguration and deletes it. Returns an error if one occurs.
func (c *dynamicBindConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dynamicBindConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dynamicBindConfiguration.
func (c *dynamicBindConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DynamicBindConfiguration, err error) {
	result = &v1.DynamicBindConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dynamicbindconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
