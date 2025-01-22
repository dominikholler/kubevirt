/*
Copyright 2024 The KubeVirt Authors.

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

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeSnapshots implements VolumeSnapshotInterface
type FakeVolumeSnapshots struct {
	Fake *FakeSnapshotV1
	ns   string
}

var volumesnapshotsResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1", Resource: "volumesnapshots"}

var volumesnapshotsKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1", Kind: "VolumeSnapshot"}

// Get takes name of the volumeSnapshot, and returns the corresponding volumeSnapshot object, and an error if there is any.
func (c *FakeVolumeSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *volumesnapshotv1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotsResource, c.ns, name), &volumesnapshotv1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshot), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshots that match those selectors.
func (c *FakeVolumeSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *volumesnapshotv1.VolumeSnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotsResource, volumesnapshotsKind, c.ns, opts), &volumesnapshotv1.VolumeSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &volumesnapshotv1.VolumeSnapshotList{ListMeta: obj.(*volumesnapshotv1.VolumeSnapshotList).ListMeta}
	for _, item := range obj.(*volumesnapshotv1.VolumeSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshots.
func (c *FakeVolumeSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotsResource, c.ns, opts))

}

// Create takes the representation of a volumeSnapshot and creates it.  Returns the server's representation of the volumeSnapshot, and an error, if there is any.
func (c *FakeVolumeSnapshots) Create(ctx context.Context, volumeSnapshot *volumesnapshotv1.VolumeSnapshot, opts v1.CreateOptions) (result *volumesnapshotv1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &volumesnapshotv1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshot), err
}

// Update takes the representation of a volumeSnapshot and updates it. Returns the server's representation of the volumeSnapshot, and an error, if there is any.
func (c *FakeVolumeSnapshots) Update(ctx context.Context, volumeSnapshot *volumesnapshotv1.VolumeSnapshot, opts v1.UpdateOptions) (result *volumesnapshotv1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &volumesnapshotv1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshots) UpdateStatus(ctx context.Context, volumeSnapshot *volumesnapshotv1.VolumeSnapshot, opts v1.UpdateOptions) (*volumesnapshotv1.VolumeSnapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotsResource, "status", c.ns, volumeSnapshot), &volumesnapshotv1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshot), err
}

// Delete takes name of the volumeSnapshot and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumesnapshotsResource, c.ns, name), &volumesnapshotv1.VolumeSnapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &volumesnapshotv1.VolumeSnapshotList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshot.
func (c *FakeVolumeSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *volumesnapshotv1.VolumeSnapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotsResource, c.ns, name, pt, data, subresources...), &volumesnapshotv1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*volumesnapshotv1.VolumeSnapshot), err
}
