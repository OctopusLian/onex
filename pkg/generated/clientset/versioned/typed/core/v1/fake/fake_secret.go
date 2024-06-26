// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecrets implements SecretInterface
type FakeSecrets struct {
	Fake *FakeCoreV1
	ns   string
}

var secretsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "secrets"}

var secretsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"}

// Get takes name of the secret, and returns the corresponding secret object, and an error if there is any.
func (c *FakeSecrets) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretsResource, c.ns, name), &corev1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Secret), err
}

// List takes label and field selectors, and returns the list of Secrets that match those selectors.
func (c *FakeSecrets) List(ctx context.Context, opts v1.ListOptions) (result *corev1.SecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretsResource, secretsKind, c.ns, opts), &corev1.SecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.SecretList{ListMeta: obj.(*corev1.SecretList).ListMeta}
	for _, item := range obj.(*corev1.SecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secrets.
func (c *FakeSecrets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretsResource, c.ns, opts))

}

// Create takes the representation of a secret and creates it.  Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Create(ctx context.Context, secret *corev1.Secret, opts v1.CreateOptions) (result *corev1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretsResource, c.ns, secret), &corev1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Secret), err
}

// Update takes the representation of a secret and updates it. Returns the server's representation of the secret, and an error, if there is any.
func (c *FakeSecrets) Update(ctx context.Context, secret *corev1.Secret, opts v1.UpdateOptions) (result *corev1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretsResource, c.ns, secret), &corev1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Secret), err
}

// Delete takes name of the secret and deletes it. Returns an error if one occurs.
func (c *FakeSecrets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(secretsResource, c.ns, name, opts), &corev1.Secret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecrets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.SecretList{})
	return err
}

// Patch applies the patch and returns the patched secret.
func (c *FakeSecrets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.Secret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretsResource, c.ns, name, pt, data, subresources...), &corev1.Secret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.Secret), err
}
