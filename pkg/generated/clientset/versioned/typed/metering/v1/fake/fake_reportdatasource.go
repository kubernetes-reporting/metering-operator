// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	meteringv1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReportDataSources implements ReportDataSourceInterface
type FakeReportDataSources struct {
	Fake *FakeMeteringV1
	ns   string
}

var reportdatasourcesResource = schema.GroupVersionResource{Group: "metering.openshift.io", Version: "v1", Resource: "reportdatasources"}

var reportdatasourcesKind = schema.GroupVersionKind{Group: "metering.openshift.io", Version: "v1", Kind: "ReportDataSource"}

// Get takes name of the reportDataSource, and returns the corresponding reportDataSource object, and an error if there is any.
func (c *FakeReportDataSources) Get(name string, options v1.GetOptions) (result *meteringv1.ReportDataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(reportdatasourcesResource, c.ns, name), &meteringv1.ReportDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportDataSource), err
}

// List takes label and ***REMOVED***eld selectors, and returns the list of ReportDataSources that match those selectors.
func (c *FakeReportDataSources) List(opts v1.ListOptions) (result *meteringv1.ReportDataSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(reportdatasourcesResource, reportdatasourcesKind, c.ns, opts), &meteringv1.ReportDataSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &meteringv1.ReportDataSourceList{ListMeta: obj.(*meteringv1.ReportDataSourceList).ListMeta}
	for _, item := range obj.(*meteringv1.ReportDataSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested reportDataSources.
func (c *FakeReportDataSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(reportdatasourcesResource, c.ns, opts))

}

// Create takes the representation of a reportDataSource and creates it.  Returns the server's representation of the reportDataSource, and an error, if there is any.
func (c *FakeReportDataSources) Create(reportDataSource *meteringv1.ReportDataSource) (result *meteringv1.ReportDataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(reportdatasourcesResource, c.ns, reportDataSource), &meteringv1.ReportDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportDataSource), err
}

// Update takes the representation of a reportDataSource and updates it. Returns the server's representation of the reportDataSource, and an error, if there is any.
func (c *FakeReportDataSources) Update(reportDataSource *meteringv1.ReportDataSource) (result *meteringv1.ReportDataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(reportdatasourcesResource, c.ns, reportDataSource), &meteringv1.ReportDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportDataSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReportDataSources) UpdateStatus(reportDataSource *meteringv1.ReportDataSource) (*meteringv1.ReportDataSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(reportdatasourcesResource, "status", c.ns, reportDataSource), &meteringv1.ReportDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportDataSource), err
}

// Delete takes name of the reportDataSource and deletes it. Returns an error if one occurs.
func (c *FakeReportDataSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(reportdatasourcesResource, c.ns, name), &meteringv1.ReportDataSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReportDataSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(reportdatasourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &meteringv1.ReportDataSourceList{})
	return err
}

// Patch applies the patch and returns the patched reportDataSource.
func (c *FakeReportDataSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *meteringv1.ReportDataSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(reportdatasourcesResource, c.ns, name, pt, data, subresources...), &meteringv1.ReportDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meteringv1.ReportDataSource), err
}