// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1"
	"github.com/operator-framework/operator-metering/pkg/generated/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MeteringV1Interface interface {
	RESTClient() rest.Interface
	HiveTablesGetter
	PrestoTablesGetter
	ReportsGetter
	ReportDataSourcesGetter
	ReportQueriesGetter
	StorageLocationsGetter
}

// MeteringV1Client is used to interact with features provided by the metering.openshift.io group.
type MeteringV1Client struct {
	restClient rest.Interface
}

func (c *MeteringV1Client) HiveTables(namespace string) HiveTableInterface {
	return newHiveTables(c, namespace)
}

func (c *MeteringV1Client) PrestoTables(namespace string) PrestoTableInterface {
	return newPrestoTables(c, namespace)
}

func (c *MeteringV1Client) Reports(namespace string) ReportInterface {
	return newReports(c, namespace)
}

func (c *MeteringV1Client) ReportDataSources(namespace string) ReportDataSourceInterface {
	return newReportDataSources(c, namespace)
}

func (c *MeteringV1Client) ReportQueries(namespace string) ReportQueryInterface {
	return newReportQueries(c, namespace)
}

func (c *MeteringV1Client) StorageLocations(namespace string) StorageLocationInterface {
	return newStorageLocations(c, namespace)
}

// NewForCon***REMOVED***g creates a new MeteringV1Client for the given con***REMOVED***g.
func NewForCon***REMOVED***g(c *rest.Con***REMOVED***g) (*MeteringV1Client, error) {
	con***REMOVED***g := *c
	if err := setCon***REMOVED***gDefaults(&con***REMOVED***g); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&con***REMOVED***g)
	if err != nil {
		return nil, err
	}
	return &MeteringV1Client{client}, nil
}

// NewForCon***REMOVED***gOrDie creates a new MeteringV1Client for the given con***REMOVED***g and
// panics if there is an error in the con***REMOVED***g.
func NewForCon***REMOVED***gOrDie(c *rest.Con***REMOVED***g) *MeteringV1Client {
	client, err := NewForCon***REMOVED***g(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MeteringV1Client for the given RESTClient.
func New(c rest.Interface) *MeteringV1Client {
	return &MeteringV1Client{c}
}

func setCon***REMOVED***gDefaults(con***REMOVED***g *rest.Con***REMOVED***g) error {
	gv := v1.SchemeGroupVersion
	con***REMOVED***g.GroupVersion = &gv
	con***REMOVED***g.APIPath = "/apis"
	con***REMOVED***g.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if con***REMOVED***g.UserAgent == "" {
		con***REMOVED***g.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MeteringV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}