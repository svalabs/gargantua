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

package v2

import (
	context "context"

	hobbyfarmiov2 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v2"
	scheme "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// UsersGetter has a method to return a UserInterface.
// A group's client should implement this interface.
type UsersGetter interface {
	Users(namespace string) UserInterface
}

// UserInterface has methods to work with User resources.
type UserInterface interface {
	Create(ctx context.Context, user *hobbyfarmiov2.User, opts v1.CreateOptions) (*hobbyfarmiov2.User, error)
	Update(ctx context.Context, user *hobbyfarmiov2.User, opts v1.UpdateOptions) (*hobbyfarmiov2.User, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*hobbyfarmiov2.User, error)
	List(ctx context.Context, opts v1.ListOptions) (*hobbyfarmiov2.UserList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hobbyfarmiov2.User, err error)
	UserExpansion
}

// users implements UserInterface
type users struct {
	*gentype.ClientWithList[*hobbyfarmiov2.User, *hobbyfarmiov2.UserList]
}

// newUsers returns a Users
func newUsers(c *HobbyfarmV2Client, namespace string) *users {
	return &users{
		gentype.NewClientWithList[*hobbyfarmiov2.User, *hobbyfarmiov2.UserList](
			"users",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *hobbyfarmiov2.User { return &hobbyfarmiov2.User{} },
			func() *hobbyfarmiov2.UserList { return &hobbyfarmiov2.UserList{} },
		),
	}
}
