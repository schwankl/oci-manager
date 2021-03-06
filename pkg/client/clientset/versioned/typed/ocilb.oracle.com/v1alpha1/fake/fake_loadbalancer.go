/*
Copyright 2018 Oracle and/or its affiliates. All rights reserved.

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
package fake

import (
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocilb.oracle.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLoadBalancers implements LoadBalancerInterface
type FakeLoadBalancers struct {
	Fake *FakeOcilbV1alpha1
	ns   string
}

var loadbalancersResource = schema.GroupVersionResource{Group: "ocilb.oracle.com", Version: "v1alpha1", Resource: "loadbalancers"}

var loadbalancersKind = schema.GroupVersionKind{Group: "ocilb.oracle.com", Version: "v1alpha1", Kind: "LoadBalancer"}

// Get takes name of the loadBalancer, and returns the corresponding loadBalancer object, and an error if there is any.
func (c *FakeLoadBalancers) Get(name string, options v1.GetOptions) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loadbalancersResource, c.ns, name), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// List takes label and field selectors, and returns the list of LoadBalancers that match those selectors.
func (c *FakeLoadBalancers) List(opts v1.ListOptions) (result *v1alpha1.LoadBalancerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loadbalancersResource, loadbalancersKind, c.ns, opts), &v1alpha1.LoadBalancerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoadBalancerList{ListMeta: obj.(*v1alpha1.LoadBalancerList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoadBalancerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loadBalancers.
func (c *FakeLoadBalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loadbalancersResource, c.ns, opts))

}

// Create takes the representation of a loadBalancer and creates it.  Returns the server's representation of the loadBalancer, and an error, if there is any.
func (c *FakeLoadBalancers) Create(loadBalancer *v1alpha1.LoadBalancer) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loadbalancersResource, c.ns, loadBalancer), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// Update takes the representation of a loadBalancer and updates it. Returns the server's representation of the loadBalancer, and an error, if there is any.
func (c *FakeLoadBalancers) Update(loadBalancer *v1alpha1.LoadBalancer) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loadbalancersResource, c.ns, loadBalancer), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}

// Delete takes name of the loadBalancer and deletes it. Returns an error if one occurs.
func (c *FakeLoadBalancers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loadbalancersResource, c.ns, name), &v1alpha1.LoadBalancer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoadBalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loadbalancersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoadBalancerList{})
	return err
}

// Patch applies the patch and returns the patched loadBalancer.
func (c *FakeLoadBalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoadBalancer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loadbalancersResource, c.ns, name, data, subresources...), &v1alpha1.LoadBalancer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancer), err
}
