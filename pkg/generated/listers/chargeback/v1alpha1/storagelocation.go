// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageLocationLister helps list StorageLocations.
type StorageLocationLister interface {
	// List lists all StorageLocations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageLocation, err error)
	// StorageLocations returns an object that can list and get StorageLocations.
	StorageLocations(namespace string) StorageLocationNamespaceLister
	StorageLocationListerExpansion
}

// storageLocationLister implements the StorageLocationLister interface.
type storageLocationLister struct {
	indexer cache.Indexer
}

// NewStorageLocationLister returns a new StorageLocationLister.
func NewStorageLocationLister(indexer cache.Indexer) StorageLocationLister {
	return &storageLocationLister{indexer: indexer}
}

// List lists all StorageLocations in the indexer.
func (s *storageLocationLister) List(selector labels.Selector) (ret []*v1alpha1.StorageLocation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageLocation))
	})
	return ret, err
}

// StorageLocations returns an object that can list and get StorageLocations.
func (s *storageLocationLister) StorageLocations(namespace string) StorageLocationNamespaceLister {
	return storageLocationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageLocationNamespaceLister helps list and get StorageLocations.
type StorageLocationNamespaceLister interface {
	// List lists all StorageLocations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageLocation, err error)
	// Get retrieves the StorageLocation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageLocation, error)
	StorageLocationNamespaceListerExpansion
}

// storageLocationNamespaceLister implements the StorageLocationNamespaceLister
// interface.
type storageLocationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageLocations in the indexer for a given namespace.
func (s storageLocationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageLocation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageLocation))
	})
	return ret, err
}

// Get retrieves the StorageLocation from the indexer for a given namespace and name.
func (s storageLocationNamespaceLister) Get(name string) (*v1alpha1.StorageLocation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagelocation"), name)
	}
	return obj.(*v1alpha1.StorageLocation), nil
}
