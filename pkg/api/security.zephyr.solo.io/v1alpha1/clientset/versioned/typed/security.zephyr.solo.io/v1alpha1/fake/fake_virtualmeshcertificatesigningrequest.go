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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMeshCertificateSigningRequests implements VirtualMeshCertificateSigningRequestInterface
type FakeVirtualMeshCertificateSigningRequests struct {
	Fake *FakeSecurityV1alpha1
	ns   string
}

var virtualmeshcertificatesigningrequestsResource = schema.GroupVersionResource{Group: "security.zephyr.solo.io", Version: "v1alpha1", Resource: "virtualmeshcertificatesigningrequests"}

var virtualmeshcertificatesigningrequestsKind = schema.GroupVersionKind{Group: "security.zephyr.solo.io", Version: "v1alpha1", Kind: "VirtualMeshCertificateSigningRequest"}

// Get takes name of the virtualMeshCertificateSigningRequest, and returns the corresponding virtualMeshCertificateSigningRequest object, and an error if there is any.
func (c *FakeVirtualMeshCertificateSigningRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMeshCertificateSigningRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmeshcertificatesigningrequestsResource, c.ns, name), &v1alpha1.VirtualMeshCertificateSigningRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMeshCertificateSigningRequest), err
}

// List takes label and field selectors, and returns the list of VirtualMeshCertificateSigningRequests that match those selectors.
func (c *FakeVirtualMeshCertificateSigningRequests) List(opts v1.ListOptions) (result *v1alpha1.VirtualMeshCertificateSigningRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmeshcertificatesigningrequestsResource, virtualmeshcertificatesigningrequestsKind, c.ns, opts), &v1alpha1.VirtualMeshCertificateSigningRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMeshCertificateSigningRequestList{ListMeta: obj.(*v1alpha1.VirtualMeshCertificateSigningRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMeshCertificateSigningRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMeshCertificateSigningRequests.
func (c *FakeVirtualMeshCertificateSigningRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmeshcertificatesigningrequestsResource, c.ns, opts))

}

// Create takes the representation of a virtualMeshCertificateSigningRequest and creates it.  Returns the server's representation of the virtualMeshCertificateSigningRequest, and an error, if there is any.
func (c *FakeVirtualMeshCertificateSigningRequests) Create(virtualMeshCertificateSigningRequest *v1alpha1.VirtualMeshCertificateSigningRequest) (result *v1alpha1.VirtualMeshCertificateSigningRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmeshcertificatesigningrequestsResource, c.ns, virtualMeshCertificateSigningRequest), &v1alpha1.VirtualMeshCertificateSigningRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMeshCertificateSigningRequest), err
}

// Update takes the representation of a virtualMeshCertificateSigningRequest and updates it. Returns the server's representation of the virtualMeshCertificateSigningRequest, and an error, if there is any.
func (c *FakeVirtualMeshCertificateSigningRequests) Update(virtualMeshCertificateSigningRequest *v1alpha1.VirtualMeshCertificateSigningRequest) (result *v1alpha1.VirtualMeshCertificateSigningRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmeshcertificatesigningrequestsResource, c.ns, virtualMeshCertificateSigningRequest), &v1alpha1.VirtualMeshCertificateSigningRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMeshCertificateSigningRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMeshCertificateSigningRequests) UpdateStatus(virtualMeshCertificateSigningRequest *v1alpha1.VirtualMeshCertificateSigningRequest) (*v1alpha1.VirtualMeshCertificateSigningRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmeshcertificatesigningrequestsResource, "status", c.ns, virtualMeshCertificateSigningRequest), &v1alpha1.VirtualMeshCertificateSigningRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMeshCertificateSigningRequest), err
}

// Delete takes name of the virtualMeshCertificateSigningRequest and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMeshCertificateSigningRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmeshcertificatesigningrequestsResource, c.ns, name), &v1alpha1.VirtualMeshCertificateSigningRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMeshCertificateSigningRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmeshcertificatesigningrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMeshCertificateSigningRequestList{})
	return err
}

// Patch applies the patch and returns the patched virtualMeshCertificateSigningRequest.
func (c *FakeVirtualMeshCertificateSigningRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMeshCertificateSigningRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmeshcertificatesigningrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VirtualMeshCertificateSigningRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMeshCertificateSigningRequest), err
}
