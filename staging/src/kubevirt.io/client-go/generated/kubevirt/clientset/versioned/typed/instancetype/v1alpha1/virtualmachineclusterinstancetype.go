/*
Copyright 2025 The KubeVirt Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevirt.io/api/instancetype/v1alpha1"
	scheme "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/scheme"
)

// VirtualMachineClusterInstancetypesGetter has a method to return a VirtualMachineClusterInstancetypeInterface.
// A group's client should implement this interface.
type VirtualMachineClusterInstancetypesGetter interface {
	VirtualMachineClusterInstancetypes() VirtualMachineClusterInstancetypeInterface
}

// VirtualMachineClusterInstancetypeInterface has methods to work with VirtualMachineClusterInstancetype resources.
type VirtualMachineClusterInstancetypeInterface interface {
	Create(ctx context.Context, virtualMachineClusterInstancetype *v1alpha1.VirtualMachineClusterInstancetype, opts v1.CreateOptions) (*v1alpha1.VirtualMachineClusterInstancetype, error)
	Update(ctx context.Context, virtualMachineClusterInstancetype *v1alpha1.VirtualMachineClusterInstancetype, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineClusterInstancetype, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VirtualMachineClusterInstancetype, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VirtualMachineClusterInstancetypeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineClusterInstancetype, err error)
	VirtualMachineClusterInstancetypeExpansion
}

// virtualMachineClusterInstancetypes implements VirtualMachineClusterInstancetypeInterface
type virtualMachineClusterInstancetypes struct {
	client rest.Interface
}

// newVirtualMachineClusterInstancetypes returns a VirtualMachineClusterInstancetypes
func newVirtualMachineClusterInstancetypes(c *InstancetypeV1alpha1Client) *virtualMachineClusterInstancetypes {
	return &virtualMachineClusterInstancetypes{
		client: c.RESTClient(),
	}
}

// Get takes name of the virtualMachineClusterInstancetype, and returns the corresponding virtualMachineClusterInstancetype object, and an error if there is any.
func (c *virtualMachineClusterInstancetypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineClusterInstancetype, err error) {
	result = &v1alpha1.VirtualMachineClusterInstancetype{}
	err = c.client.Get().
		Resource("virtualmachineclusterinstancetypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineClusterInstancetypes that match those selectors.
func (c *virtualMachineClusterInstancetypes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineClusterInstancetypeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineClusterInstancetypeList{}
	err = c.client.Get().
		Resource("virtualmachineclusterinstancetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineClusterInstancetypes.
func (c *virtualMachineClusterInstancetypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("virtualmachineclusterinstancetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineClusterInstancetype and creates it.  Returns the server's representation of the virtualMachineClusterInstancetype, and an error, if there is any.
func (c *virtualMachineClusterInstancetypes) Create(ctx context.Context, virtualMachineClusterInstancetype *v1alpha1.VirtualMachineClusterInstancetype, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineClusterInstancetype, err error) {
	result = &v1alpha1.VirtualMachineClusterInstancetype{}
	err = c.client.Post().
		Resource("virtualmachineclusterinstancetypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineClusterInstancetype).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineClusterInstancetype and updates it. Returns the server's representation of the virtualMachineClusterInstancetype, and an error, if there is any.
func (c *virtualMachineClusterInstancetypes) Update(ctx context.Context, virtualMachineClusterInstancetype *v1alpha1.VirtualMachineClusterInstancetype, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineClusterInstancetype, err error) {
	result = &v1alpha1.VirtualMachineClusterInstancetype{}
	err = c.client.Put().
		Resource("virtualmachineclusterinstancetypes").
		Name(virtualMachineClusterInstancetype.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineClusterInstancetype).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineClusterInstancetype and deletes it. Returns an error if one occurs.
func (c *virtualMachineClusterInstancetypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("virtualmachineclusterinstancetypes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineClusterInstancetypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("virtualmachineclusterinstancetypes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineClusterInstancetype.
func (c *virtualMachineClusterInstancetypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineClusterInstancetype, err error) {
	result = &v1alpha1.VirtualMachineClusterInstancetype{}
	err = c.client.Patch(pt).
		Resource("virtualmachineclusterinstancetypes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
