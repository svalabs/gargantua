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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apishobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	versioned "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/hobbyfarm/gargantua/v3/pkg/client/informers/externalversions/internalinterfaces"
	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/client/listers/hobbyfarm.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// QuizEvaluationInformer provides access to a shared informer and lister for
// QuizEvaluations.
type QuizEvaluationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() hobbyfarmiov1.QuizEvaluationLister
}

type quizEvaluationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewQuizEvaluationInformer constructs a new informer for QuizEvaluation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewQuizEvaluationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredQuizEvaluationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredQuizEvaluationInformer constructs a new informer for QuizEvaluation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredQuizEvaluationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HobbyfarmV1().QuizEvaluations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HobbyfarmV1().QuizEvaluations(namespace).Watch(context.TODO(), options)
			},
		},
		&apishobbyfarmiov1.QuizEvaluation{},
		resyncPeriod,
		indexers,
	)
}

func (f *quizEvaluationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredQuizEvaluationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *quizEvaluationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apishobbyfarmiov1.QuizEvaluation{}, f.defaultInformer)
}

func (f *quizEvaluationInformer) Lister() hobbyfarmiov1.QuizEvaluationLister {
	return hobbyfarmiov1.NewQuizEvaluationLister(f.Informer().GetIndexer())
}
