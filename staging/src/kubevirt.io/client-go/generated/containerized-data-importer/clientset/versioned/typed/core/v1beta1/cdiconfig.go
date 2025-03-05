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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned/scheme"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// CDIConfigsGetter has a method to return a CDIConfigInterface.
// A group's client should implement this interface.
type CDIConfigsGetter interface {
	CDIConfigs() CDIConfigInterface
}

// CDIConfigInterface has methods to work with CDIConfig resources.
type CDIConfigInterface interface {
	Create(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.CreateOptions) (*v1beta1.CDIConfig, error)
	Update(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.UpdateOptions) (*v1beta1.CDIConfig, error)
	UpdateStatus(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.UpdateOptions) (*v1beta1.CDIConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.CDIConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CDIConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CDIConfig, err error)
	CDIConfigExpansion
}

// cDIConfigs implements CDIConfigInterface
type cDIConfigs struct {
	client rest.Interface
}

// newCDIConfigs returns a CDIConfigs
func newCDIConfigs(c *CdiV1beta1Client) *cDIConfigs {
	return &cDIConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the cDIConfig, and returns the corresponding cDIConfig object, and an error if there is any.
func (c *cDIConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CDIConfig, err error) {
	result = &v1beta1.CDIConfig{}
	err = c.client.Get().
		Resource("cdiconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CDIConfigs that match those selectors.
func (c *cDIConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CDIConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CDIConfigList{}
	err = c.client.Get().
		Resource("cdiconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cDIConfigs.
func (c *cDIConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("cdiconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cDIConfig and creates it.  Returns the server's representation of the cDIConfig, and an error, if there is any.
func (c *cDIConfigs) Create(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.CreateOptions) (result *v1beta1.CDIConfig, err error) {
	result = &v1beta1.CDIConfig{}
	err = c.client.Post().
		Resource("cdiconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cDIConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cDIConfig and updates it. Returns the server's representation of the cDIConfig, and an error, if there is any.
func (c *cDIConfigs) Update(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.UpdateOptions) (result *v1beta1.CDIConfig, err error) {
	result = &v1beta1.CDIConfig{}
	err = c.client.Put().
		Resource("cdiconfigs").
		Name(cDIConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cDIConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cDIConfigs) UpdateStatus(ctx context.Context, cDIConfig *v1beta1.CDIConfig, opts v1.UpdateOptions) (result *v1beta1.CDIConfig, err error) {
	result = &v1beta1.CDIConfig{}
	err = c.client.Put().
		Resource("cdiconfigs").
		Name(cDIConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cDIConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cDIConfig and deletes it. Returns an error if one occurs.
func (c *cDIConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("cdiconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cDIConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("cdiconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cDIConfig.
func (c *cDIConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CDIConfig, err error) {
	result = &v1beta1.CDIConfig{}
	err = c.client.Patch(pt).
		Resource("cdiconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
