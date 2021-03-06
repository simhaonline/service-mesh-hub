package kube

import (
	"bytes"

	"github.com/hashicorp/go-multierror"
	"k8s.io/cli-runtime/pkg/resource"
)

//go:generate mockgen -source ./unstructured_kube_client.go -destination mocks/mock_unstructured_kube_client.go
type UnstructuredKubeClient interface {
	// build unstructured k8s objects out of the string representation of a manifest
	BuildResources(namespace string, manifest string) ([]*resource.Info, error)

	// create as many resources as we can, returning the ones that were successful along with any error that occurred
	Create(namespace string, resources []*resource.Info) (createdResources []*resource.Info, err error)

	// delete as many of the given resources as we can
	Delete(namespace string, resources []*resource.Info) (deletedResources []*resource.Info, err error)
}

type UnstructuredKubeClientFactory func(restClientGetter resource.RESTClientGetter) UnstructuredKubeClient

func NewUnstructuredKubeClientFactory() UnstructuredKubeClientFactory {
	return func(restClientGetter resource.RESTClientGetter) UnstructuredKubeClient {
		return &unstructuredKubeClient{
			restClientGetter: restClientGetter,
		}
	}
}

type unstructuredKubeClient struct {
	restClientGetter resource.RESTClientGetter
}

func (u *unstructuredKubeClient) BuildResources(namespace string, manifest string) ([]*resource.Info, error) {
	return resource.NewBuilder(u.restClientGetter).
		RequireNamespace().
		NamespaceParam(namespace).
		DefaultNamespace().
		Unstructured().
		Stream(bytes.NewBuffer([]byte(manifest)), "").
		Do().
		Infos()
}

func (u *unstructuredKubeClient) Create(namespace string, resources []*resource.Info) (createdResources []*resource.Info, err error) {
	var multiErr *multierror.Error
	for _, r := range resources {
		_, err := resource.NewHelper(r.Client, r.Mapping).Create(namespace, false, r.Object, nil)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
			continue
		} else {
			createdResources = append(createdResources, r)
		}
	}

	// Note: Go has typed nils, so doing just `return multiErr` will result in a panic due to dereferenced nil
	// I discovered this after painstaking debugging
	return createdResources, multiErr.ErrorOrNil()
}

func (u *unstructuredKubeClient) Delete(namespace string, resources []*resource.Info) (deletedResources []*resource.Info, err error) {
	for _, r := range resources {
		_, err := resource.NewHelper(r.Client, r.Mapping).Delete(namespace, r.Name)
		if err != nil {
			return deletedResources, err
		}
		deletedResources = append(deletedResources, r)
	}

	return deletedResources, nil
}
