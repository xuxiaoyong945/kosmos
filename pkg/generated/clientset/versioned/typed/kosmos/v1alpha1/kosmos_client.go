// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/kosmos.io/kosmos/pkg/apis/kosmos/v1alpha1"
	"github.com/kosmos.io/kosmos/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type KosmosV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClustersGetter
	ClusterDistributionPoliciesGetter
	ClusterNodesGetter
	ClusterPodConvertPoliciesGetter
	DaemonSetsGetter
	DistributionPoliciesGetter
	KnodesGetter
	NodeConfigsGetter
	PodConvertPoliciesGetter
	ShadowDaemonSetsGetter
	VirtualClustersGetter
}

// KosmosV1alpha1Client is used to interact with features provided by the kosmos.io group.
type KosmosV1alpha1Client struct {
	restClient rest.Interface
}

func (c *KosmosV1alpha1Client) Clusters() ClusterInterface {
	return newClusters(c)
}

func (c *KosmosV1alpha1Client) ClusterDistributionPolicies() ClusterDistributionPolicyInterface {
	return newClusterDistributionPolicies(c)
}

func (c *KosmosV1alpha1Client) ClusterNodes() ClusterNodeInterface {
	return newClusterNodes(c)
}

func (c *KosmosV1alpha1Client) ClusterPodConvertPolicies() ClusterPodConvertPolicyInterface {
	return newClusterPodConvertPolicies(c)
}

func (c *KosmosV1alpha1Client) DaemonSets(namespace string) DaemonSetInterface {
	return newDaemonSets(c, namespace)
}

func (c *KosmosV1alpha1Client) DistributionPolicies(namespace string) DistributionPolicyInterface {
	return newDistributionPolicies(c, namespace)
}

func (c *KosmosV1alpha1Client) Knodes() KnodeInterface {
	return newKnodes(c)
}

func (c *KosmosV1alpha1Client) NodeConfigs() NodeConfigInterface {
	return newNodeConfigs(c)
}

func (c *KosmosV1alpha1Client) PodConvertPolicies(namespace string) PodConvertPolicyInterface {
	return newPodConvertPolicies(c, namespace)
}

func (c *KosmosV1alpha1Client) ShadowDaemonSets(namespace string) ShadowDaemonSetInterface {
	return newShadowDaemonSets(c, namespace)
}

func (c *KosmosV1alpha1Client) VirtualClusters() VirtualClusterInterface {
	return newVirtualClusters(c)
}

// NewForConfig creates a new KosmosV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KosmosV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KosmosV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KosmosV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KosmosV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new KosmosV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KosmosV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KosmosV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *KosmosV1alpha1Client {
	return &KosmosV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KosmosV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
