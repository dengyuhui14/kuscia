// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainDataGrantLister helps list DomainDataGrants.
// All objects returned here must be treated as read-only.
type DomainDataGrantLister interface {
	// List lists all DomainDataGrants in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainDataGrant, err error)
	// DomainDataGrants returns an object that can list and get DomainDataGrants.
	DomainDataGrants(namespace string) DomainDataGrantNamespaceLister
	DomainDataGrantListerExpansion
}

// domainDataGrantLister implements the DomainDataGrantLister interface.
type domainDataGrantLister struct {
	indexer cache.Indexer
}

// NewDomainDataGrantLister returns a new DomainDataGrantLister.
func NewDomainDataGrantLister(indexer cache.Indexer) DomainDataGrantLister {
	return &domainDataGrantLister{indexer: indexer}
}

// List lists all DomainDataGrants in the indexer.
func (s *domainDataGrantLister) List(selector labels.Selector) (ret []*v1alpha1.DomainDataGrant, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainDataGrant))
	})
	return ret, err
}

// DomainDataGrants returns an object that can list and get DomainDataGrants.
func (s *domainDataGrantLister) DomainDataGrants(namespace string) DomainDataGrantNamespaceLister {
	return domainDataGrantNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainDataGrantNamespaceLister helps list and get DomainDataGrants.
// All objects returned here must be treated as read-only.
type DomainDataGrantNamespaceLister interface {
	// List lists all DomainDataGrants in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainDataGrant, err error)
	// Get retrieves the DomainDataGrant from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainDataGrant, error)
	DomainDataGrantNamespaceListerExpansion
}

// domainDataGrantNamespaceLister implements the DomainDataGrantNamespaceLister
// interface.
type domainDataGrantNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainDataGrants in the indexer for a given namespace.
func (s domainDataGrantNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainDataGrant, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainDataGrant))
	})
	return ret, err
}

// Get retrieves the DomainDataGrant from the indexer for a given namespace and name.
func (s domainDataGrantNamespaceLister) Get(name string) (*v1alpha1.DomainDataGrant, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domaindatagrant"), name)
	}
	return obj.(*v1alpha1.DomainDataGrant), nil
}
