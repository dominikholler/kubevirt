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

package fake

import (
	"context"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeThanosRulers implements ThanosRulerInterface
type FakeThanosRulers struct {
	Fake *FakeMonitoringV1
	ns   string
}

var thanosrulersResource = schema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1", Resource: "thanosrulers"}

var thanosrulersKind = schema.GroupVersionKind{Group: "monitoring.coreos.com", Version: "v1", Kind: "ThanosRuler"}

// Get takes name of the thanosRuler, and returns the corresponding thanosRuler object, and an error if there is any.
func (c *FakeThanosRulers) Get(ctx context.Context, name string, options v1.GetOptions) (result *monitoringv1.ThanosRuler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(thanosrulersResource, c.ns, name), &monitoringv1.ThanosRuler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ThanosRuler), err
}

// List takes label and field selectors, and returns the list of ThanosRulers that match those selectors.
func (c *FakeThanosRulers) List(ctx context.Context, opts v1.ListOptions) (result *monitoringv1.ThanosRulerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(thanosrulersResource, thanosrulersKind, c.ns, opts), &monitoringv1.ThanosRulerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &monitoringv1.ThanosRulerList{ListMeta: obj.(*monitoringv1.ThanosRulerList).ListMeta}
	for _, item := range obj.(*monitoringv1.ThanosRulerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested thanosRulers.
func (c *FakeThanosRulers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(thanosrulersResource, c.ns, opts))

}

// Create takes the representation of a thanosRuler and creates it.  Returns the server's representation of the thanosRuler, and an error, if there is any.
func (c *FakeThanosRulers) Create(ctx context.Context, thanosRuler *monitoringv1.ThanosRuler, opts v1.CreateOptions) (result *monitoringv1.ThanosRuler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(thanosrulersResource, c.ns, thanosRuler), &monitoringv1.ThanosRuler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ThanosRuler), err
}

// Update takes the representation of a thanosRuler and updates it. Returns the server's representation of the thanosRuler, and an error, if there is any.
func (c *FakeThanosRulers) Update(ctx context.Context, thanosRuler *monitoringv1.ThanosRuler, opts v1.UpdateOptions) (result *monitoringv1.ThanosRuler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(thanosrulersResource, c.ns, thanosRuler), &monitoringv1.ThanosRuler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ThanosRuler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeThanosRulers) UpdateStatus(ctx context.Context, thanosRuler *monitoringv1.ThanosRuler, opts v1.UpdateOptions) (*monitoringv1.ThanosRuler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(thanosrulersResource, "status", c.ns, thanosRuler), &monitoringv1.ThanosRuler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ThanosRuler), err
}

// Delete takes name of the thanosRuler and deletes it. Returns an error if one occurs.
func (c *FakeThanosRulers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(thanosrulersResource, c.ns, name), &monitoringv1.ThanosRuler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeThanosRulers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(thanosrulersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &monitoringv1.ThanosRulerList{})
	return err
}

// Patch applies the patch and returns the patched thanosRuler.
func (c *FakeThanosRulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *monitoringv1.ThanosRuler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(thanosrulersResource, c.ns, name, pt, data, subresources...), &monitoringv1.ThanosRuler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*monitoringv1.ThanosRuler), err
}
