/*
Copyright 2020 The Tilt Dev Authors

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

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	corev1alpha1 "github.com/tilt-dev/tilt/pkg/apis/core/v1alpha1"
	versioned "github.com/tilt-dev/tilt-api-client-go/pkg/clientset/versioned"
	internalinterfaces "github.com/tilt-dev/tilt-api-client-go/pkg/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tilt-dev/tilt-api-client-go/pkg/listers/core/v1alpha1"
)

// SessionInformer provides access to a shared informer and lister for
// Sessions.
type SessionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SessionLister
}

type sessionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSessionInformer constructs a new informer for Session type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSessionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSessionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSessionInformer constructs a new informer for Session type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSessionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TiltV1alpha1().Sessions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TiltV1alpha1().Sessions().Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.Session{},
		resyncPeriod,
		indexers,
	)
}

func (f *sessionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSessionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sessionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.Session{}, f.defaultInformer)
}

func (f *sessionInformer) Lister() v1alpha1.SessionLister {
	return v1alpha1.NewSessionLister(f.Informer().GetIndexer())
}
