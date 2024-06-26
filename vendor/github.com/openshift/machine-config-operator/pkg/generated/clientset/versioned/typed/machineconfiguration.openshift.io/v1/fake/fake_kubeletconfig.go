// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	machineconfigurationopenshiftiov1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeletConfigs implements KubeletConfigInterface
type FakeKubeletConfigs struct {
	Fake *FakeMachineconfigurationV1
}

var kubeletconfigsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "kubeletconfigs"}

var kubeletconfigsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "KubeletConfig"}

// Get takes name of the kubeletConfig, and returns the corresponding kubeletConfig object, and an error if there is any.
func (c *FakeKubeletConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *machineconfigurationopenshiftiov1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(kubeletconfigsResource, name), &machineconfigurationopenshiftiov1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.KubeletConfig), err
}

// List takes label and field selectors, and returns the list of KubeletConfigs that match those selectors.
func (c *FakeKubeletConfigs) List(ctx context.Context, opts v1.ListOptions) (result *machineconfigurationopenshiftiov1.KubeletConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(kubeletconfigsResource, kubeletconfigsKind, opts), &machineconfigurationopenshiftiov1.KubeletConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationopenshiftiov1.KubeletConfigList{ListMeta: obj.(*machineconfigurationopenshiftiov1.KubeletConfigList).ListMeta}
	for _, item := range obj.(*machineconfigurationopenshiftiov1.KubeletConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeletConfigs.
func (c *FakeKubeletConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(kubeletconfigsResource, opts))
}

// Create takes the representation of a kubeletConfig and creates it.  Returns the server's representation of the kubeletConfig, and an error, if there is any.
func (c *FakeKubeletConfigs) Create(ctx context.Context, kubeletConfig *machineconfigurationopenshiftiov1.KubeletConfig, opts v1.CreateOptions) (result *machineconfigurationopenshiftiov1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(kubeletconfigsResource, kubeletConfig), &machineconfigurationopenshiftiov1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.KubeletConfig), err
}

// Update takes the representation of a kubeletConfig and updates it. Returns the server's representation of the kubeletConfig, and an error, if there is any.
func (c *FakeKubeletConfigs) Update(ctx context.Context, kubeletConfig *machineconfigurationopenshiftiov1.KubeletConfig, opts v1.UpdateOptions) (result *machineconfigurationopenshiftiov1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(kubeletconfigsResource, kubeletConfig), &machineconfigurationopenshiftiov1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.KubeletConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeletConfigs) UpdateStatus(ctx context.Context, kubeletConfig *machineconfigurationopenshiftiov1.KubeletConfig, opts v1.UpdateOptions) (*machineconfigurationopenshiftiov1.KubeletConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(kubeletconfigsResource, "status", kubeletConfig), &machineconfigurationopenshiftiov1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.KubeletConfig), err
}

// Delete takes name of the kubeletConfig and deletes it. Returns an error if one occurs.
func (c *FakeKubeletConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(kubeletconfigsResource, name, opts), &machineconfigurationopenshiftiov1.KubeletConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeletConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(kubeletconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &machineconfigurationopenshiftiov1.KubeletConfigList{})
	return err
}

// Patch applies the patch and returns the patched kubeletConfig.
func (c *FakeKubeletConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machineconfigurationopenshiftiov1.KubeletConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(kubeletconfigsResource, name, pt, data, subresources...), &machineconfigurationopenshiftiov1.KubeletConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationopenshiftiov1.KubeletConfig), err
}
