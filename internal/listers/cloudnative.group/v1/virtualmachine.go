/*
Copyright 2022.

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
	v1 "controller/pkg/apis/cloudnative.group/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualMachineLister helps list VirtualMachines.
// All objects returned here must be treated as read-only.
type VirtualMachineLister interface {
	// List lists all VirtualMachines in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualMachine, err error)
	// VirtualMachines returns an object that can list and get VirtualMachines.
	VirtualMachines(namespace string) VirtualMachineNamespaceLister
	VirtualMachineListerExpansion
}

// virtualMachineLister implements the VirtualMachineLister interface.
type virtualMachineLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineLister returns a new VirtualMachineLister.
func NewVirtualMachineLister(indexer cache.Indexer) VirtualMachineLister {
	return &virtualMachineLister{indexer: indexer}
}

// List lists all VirtualMachines in the indexer.
func (s *virtualMachineLister) List(selector labels.Selector) (ret []*v1.VirtualMachine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualMachine))
	})
	return ret, err
}

// VirtualMachines returns an object that can list and get VirtualMachines.
func (s *virtualMachineLister) VirtualMachines(namespace string) VirtualMachineNamespaceLister {
	return virtualMachineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineNamespaceLister helps list and get VirtualMachines.
// All objects returned here must be treated as read-only.
type VirtualMachineNamespaceLister interface {
	// List lists all VirtualMachines in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualMachine, err error)
	// Get retrieves the VirtualMachine from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualMachine, error)
	VirtualMachineNamespaceListerExpansion
}

// virtualMachineNamespaceLister implements the VirtualMachineNamespaceLister
// interface.
type virtualMachineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachines in the indexer for a given namespace.
func (s virtualMachineNamespaceLister) List(selector labels.Selector) (ret []*v1.VirtualMachine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualMachine))
	})
	return ret, err
}

// Get retrieves the VirtualMachine from the indexer for a given namespace and name.
func (s virtualMachineNamespaceLister) Get(name string) (*v1.VirtualMachine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualmachine"), name)
	}
	return obj.(*v1.VirtualMachine), nil
}
