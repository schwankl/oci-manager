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
package v1alpha1

import (
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocilb.oracle.com/v1alpha1"
	"github.com/oracle/oci-manager/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type OcilbV1alpha1Interface interface {
	RESTClient() rest.Interface
	BackendsGetter
	BackendSetsGetter
	CertificatesGetter
	ListenersGetter
	LoadBalancersGetter
}

// OcilbV1alpha1Client is used to interact with features provided by the ocilb.oracle.com group.
type OcilbV1alpha1Client struct {
	restClient rest.Interface
}

func (c *OcilbV1alpha1Client) Backends(namespace string) BackendInterface {
	return newBackends(c, namespace)
}

func (c *OcilbV1alpha1Client) BackendSets(namespace string) BackendSetInterface {
	return newBackendSets(c, namespace)
}

func (c *OcilbV1alpha1Client) Certificates(namespace string) CertificateInterface {
	return newCertificates(c, namespace)
}

func (c *OcilbV1alpha1Client) Listeners(namespace string) ListenerInterface {
	return newListeners(c, namespace)
}

func (c *OcilbV1alpha1Client) LoadBalancers(namespace string) LoadBalancerInterface {
	return newLoadBalancers(c, namespace)
}

// NewForConfig creates a new OcilbV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*OcilbV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OcilbV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new OcilbV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OcilbV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OcilbV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *OcilbV1alpha1Client {
	return &OcilbV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OcilbV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
