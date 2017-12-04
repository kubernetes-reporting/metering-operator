// This file was automatically generated by informer-gen

package v1alpha1

import (
	chargeback_v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	versioned "github.com/coreos-inc/kube-chargeback/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/coreos-inc/kube-chargeback/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/generated/listers/chargeback/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// StorageLocationInformer provides access to a shared informer and lister for
// StorageLocations.
type StorageLocationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageLocationLister
}

type storageLocationInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewStorageLocationInformer constructs a new informer for StorageLocation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageLocationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ChargebackV1alpha1().StorageLocations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ChargebackV1alpha1().StorageLocations(namespace).Watch(options)
			},
		},
		&chargeback_v1alpha1.StorageLocation{},
		resyncPeriod,
		indexers,
	)
}

func defaultStorageLocationInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewStorageLocationInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *storageLocationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&chargeback_v1alpha1.StorageLocation{}, defaultStorageLocationInformer)
}

func (f *storageLocationInformer) Lister() v1alpha1.StorageLocationLister {
	return v1alpha1.NewStorageLocationLister(f.Informer().GetIndexer())
}
