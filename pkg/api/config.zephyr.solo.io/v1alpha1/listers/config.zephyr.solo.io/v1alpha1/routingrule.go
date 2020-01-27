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

package v1alpha1

import (
	v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/config.zephyr.solo.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoutingRuleLister helps list RoutingRules.
type RoutingRuleLister interface {
	// List lists all RoutingRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RoutingRule, err error)
	// RoutingRules returns an object that can list and get RoutingRules.
	RoutingRules(namespace string) RoutingRuleNamespaceLister
	RoutingRuleListerExpansion
}

// routingRuleLister implements the RoutingRuleLister interface.
type routingRuleLister struct {
	indexer cache.Indexer
}

// NewRoutingRuleLister returns a new RoutingRuleLister.
func NewRoutingRuleLister(indexer cache.Indexer) RoutingRuleLister {
	return &routingRuleLister{indexer: indexer}
}

// List lists all RoutingRules in the indexer.
func (s *routingRuleLister) List(selector labels.Selector) (ret []*v1alpha1.RoutingRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RoutingRule))
	})
	return ret, err
}

// RoutingRules returns an object that can list and get RoutingRules.
func (s *routingRuleLister) RoutingRules(namespace string) RoutingRuleNamespaceLister {
	return routingRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RoutingRuleNamespaceLister helps list and get RoutingRules.
type RoutingRuleNamespaceLister interface {
	// List lists all RoutingRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RoutingRule, err error)
	// Get retrieves the RoutingRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RoutingRule, error)
	RoutingRuleNamespaceListerExpansion
}

// routingRuleNamespaceLister implements the RoutingRuleNamespaceLister
// interface.
type routingRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RoutingRules in the indexer for a given namespace.
func (s routingRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RoutingRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RoutingRule))
	})
	return ret, err
}

// Get retrieves the RoutingRule from the indexer for a given namespace and name.
func (s routingRuleNamespaceLister) Get(name string) (*v1alpha1.RoutingRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("routingrule"), name)
	}
	return obj.(*v1alpha1.RoutingRule), nil
}