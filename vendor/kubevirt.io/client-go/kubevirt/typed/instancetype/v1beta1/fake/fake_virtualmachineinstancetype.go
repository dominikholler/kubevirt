/*
Copyright The KubeVirt Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubevirt.io/api/instancetype/v1beta1"
)

// FakeVirtualMachineInstancetypes implements VirtualMachineInstancetypeInterface
type FakeVirtualMachineInstancetypes struct {
	Fake *FakeInstancetypeV1beta1
	ns   string
}

var virtualmachineinstancetypesResource = v1beta1.SchemeGroupVersion.WithResource("virtualmachineinstancetypes")

var virtualmachineinstancetypesKind = v1beta1.SchemeGroupVersion.WithKind("VirtualMachineInstancetype")

// Get takes name of the virtualMachineInstancetype, and returns the corresponding virtualMachineInstancetype object, and an error if there is any.
func (c *FakeVirtualMachineInstancetypes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineInstancetype, err error) {
	emptyResult := &v1beta1.VirtualMachineInstancetype{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(virtualmachineinstancetypesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.VirtualMachineInstancetype), err
}

// List takes label and field selectors, and returns the list of VirtualMachineInstancetypes that match those selectors.
func (c *FakeVirtualMachineInstancetypes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineInstancetypeList, err error) {
	emptyResult := &v1beta1.VirtualMachineInstancetypeList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(virtualmachineinstancetypesResource, virtualmachineinstancetypesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineInstancetypeList{ListMeta: obj.(*v1beta1.VirtualMachineInstancetypeList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineInstancetypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineInstancetypes.
func (c *FakeVirtualMachineInstancetypes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(virtualmachineinstancetypesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineInstancetype and creates it.  Returns the server's representation of the virtualMachineInstancetype, and an error, if there is any.
func (c *FakeVirtualMachineInstancetypes) Create(ctx context.Context, virtualMachineInstancetype *v1beta1.VirtualMachineInstancetype, opts v1.CreateOptions) (result *v1beta1.VirtualMachineInstancetype, err error) {
	emptyResult := &v1beta1.VirtualMachineInstancetype{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(virtualmachineinstancetypesResource, c.ns, virtualMachineInstancetype, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.VirtualMachineInstancetype), err
}

// Update takes the representation of a virtualMachineInstancetype and updates it. Returns the server's representation of the virtualMachineInstancetype, and an error, if there is any.
func (c *FakeVirtualMachineInstancetypes) Update(ctx context.Context, virtualMachineInstancetype *v1beta1.VirtualMachineInstancetype, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineInstancetype, err error) {
	emptyResult := &v1beta1.VirtualMachineInstancetype{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(virtualmachineinstancetypesResource, c.ns, virtualMachineInstancetype, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.VirtualMachineInstancetype), err
}

// Delete takes name of the virtualMachineInstancetype and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineInstancetypes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualmachineinstancetypesResource, c.ns, name, opts), &v1beta1.VirtualMachineInstancetype{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineInstancetypes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(virtualmachineinstancetypesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineInstancetypeList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineInstancetype.
func (c *FakeVirtualMachineInstancetypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineInstancetype, err error) {
	emptyResult := &v1beta1.VirtualMachineInstancetype{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(virtualmachineinstancetypesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.VirtualMachineInstancetype), err
}
