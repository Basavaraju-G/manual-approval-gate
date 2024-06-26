/*
Copyright 2022 The OpenShift Pipelines Authors

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

	v1alpha1 "github.com/openshift-pipelines/manual-approval-gate/pkg/apis/approvaltask/v1alpha1"
	scheme "github.com/openshift-pipelines/manual-approval-gate/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApprovalTasksGetter has a method to return a ApprovalTaskInterface.
// A group's client should implement this interface.
type ApprovalTasksGetter interface {
	ApprovalTasks(namespace string) ApprovalTaskInterface
}

// ApprovalTaskInterface has methods to work with ApprovalTask resources.
type ApprovalTaskInterface interface {
	Create(ctx context.Context, approvalTask *v1alpha1.ApprovalTask, opts v1.CreateOptions) (*v1alpha1.ApprovalTask, error)
	Update(ctx context.Context, approvalTask *v1alpha1.ApprovalTask, opts v1.UpdateOptions) (*v1alpha1.ApprovalTask, error)
	UpdateStatus(ctx context.Context, approvalTask *v1alpha1.ApprovalTask, opts v1.UpdateOptions) (*v1alpha1.ApprovalTask, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApprovalTask, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApprovalTaskList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApprovalTask, err error)
	ApprovalTaskExpansion
}

// approvalTasks implements ApprovalTaskInterface
type approvalTasks struct {
	client rest.Interface
	ns     string
}

// newApprovalTasks returns a ApprovalTasks
func newApprovalTasks(c *OpenshiftpipelinesV1alpha1Client, namespace string) *approvalTasks {
	return &approvalTasks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the approvalTask, and returns the corresponding approvalTask object, and an error if there is any.
func (c *approvalTasks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApprovalTask, err error) {
	result = &v1alpha1.ApprovalTask{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("approvaltasks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApprovalTasks that match those selectors.
func (c *approvalTasks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApprovalTaskList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApprovalTaskList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("approvaltasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested approvalTasks.
func (c *approvalTasks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("approvaltasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a approvalTask and creates it.  Returns the server's representation of the approvalTask, and an error, if there is any.
func (c *approvalTasks) Create(ctx context.Context, approvalTask *v1alpha1.ApprovalTask, opts v1.CreateOptions) (result *v1alpha1.ApprovalTask, err error) {
	result = &v1alpha1.ApprovalTask{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("approvaltasks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(approvalTask).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a approvalTask and updates it. Returns the server's representation of the approvalTask, and an error, if there is any.
func (c *approvalTasks) Update(ctx context.Context, approvalTask *v1alpha1.ApprovalTask, opts v1.UpdateOptions) (result *v1alpha1.ApprovalTask, err error) {
	result = &v1alpha1.ApprovalTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("approvaltasks").
		Name(approvalTask.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(approvalTask).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *approvalTasks) UpdateStatus(ctx context.Context, approvalTask *v1alpha1.ApprovalTask, opts v1.UpdateOptions) (result *v1alpha1.ApprovalTask, err error) {
	result = &v1alpha1.ApprovalTask{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("approvaltasks").
		Name(approvalTask.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(approvalTask).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the approvalTask and deletes it. Returns an error if one occurs.
func (c *approvalTasks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("approvaltasks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *approvalTasks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("approvaltasks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched approvalTask.
func (c *approvalTasks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApprovalTask, err error) {
	result = &v1alpha1.ApprovalTask{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("approvaltasks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
