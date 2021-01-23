/*
Copyright 2020 The Knative Authors

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
	scheme "knative.dev/net-ingressv2/pkg/client/clientset/versioned/scheme"
	v1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
)

// GatewaysGetter has a method to return a GatewayInterface.
// A group's client should implement this interface.
type GatewaysGetter interface {
	Gateways(namespace string) GatewayInterface
}

// GatewayInterface has methods to work with Gateway resources.
type GatewayInterface interface {
	Create(ctx context.Context, gateway *v1alpha1.Gateway, opts v1.CreateOptions) (*v1alpha1.Gateway, error)
	Update(ctx context.Context, gateway *v1alpha1.Gateway, opts v1.UpdateOptions) (*v1alpha1.Gateway, error)
	UpdateStatus(ctx context.Context, gateway *v1alpha1.Gateway, opts v1.UpdateOptions) (*v1alpha1.Gateway, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Gateway, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GatewayList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Gateway, err error)
	GatewayExpansion
}

// gateways implements GatewayInterface
type gateways struct {
	client rest.Interface
	ns     string
}

// newGateways returns a Gateways
func newGateways(c *NetworkingV1alpha1Client, namespace string) *gateways {
	return &gateways{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gateway, and returns the corresponding gateway object, and an error if there is any.
func (c *gateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Gateway, err error) {
	result = &v1alpha1.Gateway{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gateways").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Gateways that match those selectors.
func (c *gateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GatewayList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GatewayList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gateways.
func (c *gateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gateway and creates it.  Returns the server's representation of the gateway, and an error, if there is any.
func (c *gateways) Create(ctx context.Context, gateway *v1alpha1.Gateway, opts v1.CreateOptions) (result *v1alpha1.Gateway, err error) {
	result = &v1alpha1.Gateway{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gateways").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gateway).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gateway and updates it. Returns the server's representation of the gateway, and an error, if there is any.
func (c *gateways) Update(ctx context.Context, gateway *v1alpha1.Gateway, opts v1.UpdateOptions) (result *v1alpha1.Gateway, err error) {
	result = &v1alpha1.Gateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gateways").
		Name(gateway.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gateway).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gateways) UpdateStatus(ctx context.Context, gateway *v1alpha1.Gateway, opts v1.UpdateOptions) (result *v1alpha1.Gateway, err error) {
	result = &v1alpha1.Gateway{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gateways").
		Name(gateway.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gateway).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gateway and deletes it. Returns an error if one occurs.
func (c *gateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gateways").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gateways").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gateway.
func (c *gateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Gateway, err error) {
	result = &v1alpha1.Gateway{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gateways").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}