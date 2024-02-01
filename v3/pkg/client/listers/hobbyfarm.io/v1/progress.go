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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProgressLister helps list Progresses.
// All objects returned here must be treated as read-only.
type ProgressLister interface {
	// List lists all Progresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Progress, err error)
	// Progresses returns an object that can list and get Progresses.
	Progresses(namespace string) ProgressNamespaceLister
	ProgressListerExpansion
}

// progressLister implements the ProgressLister interface.
type progressLister struct {
	indexer cache.Indexer
}

// NewProgressLister returns a new ProgressLister.
func NewProgressLister(indexer cache.Indexer) ProgressLister {
	return &progressLister{indexer: indexer}
}

// List lists all Progresses in the indexer.
func (s *progressLister) List(selector labels.Selector) (ret []*v1.Progress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Progress))
	})
	return ret, err
}

// Progresses returns an object that can list and get Progresses.
func (s *progressLister) Progresses(namespace string) ProgressNamespaceLister {
	return progressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProgressNamespaceLister helps list and get Progresses.
// All objects returned here must be treated as read-only.
type ProgressNamespaceLister interface {
	// List lists all Progresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Progress, err error)
	// Get retrieves the Progress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Progress, error)
	ProgressNamespaceListerExpansion
}

// progressNamespaceLister implements the ProgressNamespaceLister
// interface.
type progressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Progresses in the indexer for a given namespace.
func (s progressNamespaceLister) List(selector labels.Selector) (ret []*v1.Progress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Progress))
	})
	return ret, err
}

// Get retrieves the Progress from the indexer for a given namespace and name.
func (s progressNamespaceLister) Get(name string) (*v1.Progress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("progress"), name)
	}
	return obj.(*v1.Progress), nil
}