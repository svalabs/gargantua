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
	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// QuizLister helps list Quizes.
// All objects returned here must be treated as read-only.
type QuizLister interface {
	// List lists all Quizes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*hobbyfarmiov1.Quiz, err error)
	// Quizes returns an object that can list and get Quizes.
	Quizes(namespace string) QuizNamespaceLister
	QuizListerExpansion
}

// quizLister implements the QuizLister interface.
type quizLister struct {
	listers.ResourceIndexer[*hobbyfarmiov1.Quiz]
}

// NewQuizLister returns a new QuizLister.
func NewQuizLister(indexer cache.Indexer) QuizLister {
	return &quizLister{listers.New[*hobbyfarmiov1.Quiz](indexer, hobbyfarmiov1.Resource("quiz"))}
}

// Quizes returns an object that can list and get Quizes.
func (s *quizLister) Quizes(namespace string) QuizNamespaceLister {
	return quizNamespaceLister{listers.NewNamespaced[*hobbyfarmiov1.Quiz](s.ResourceIndexer, namespace)}
}

// QuizNamespaceLister helps list and get Quizes.
// All objects returned here must be treated as read-only.
type QuizNamespaceLister interface {
	// List lists all Quizes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*hobbyfarmiov1.Quiz, err error)
	// Get retrieves the Quiz from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*hobbyfarmiov1.Quiz, error)
	QuizNamespaceListerExpansion
}

// quizNamespaceLister implements the QuizNamespaceLister
// interface.
type quizNamespaceLister struct {
	listers.ResourceIndexer[*hobbyfarmiov1.Quiz]
}
