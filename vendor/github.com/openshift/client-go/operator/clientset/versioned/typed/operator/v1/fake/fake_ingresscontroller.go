// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIngressControllers implements IngressControllerInterface
type FakeIngressControllers struct {
	Fake *FakeOperatorV1
	ns   string
}

var ingresscontrollersResource = schema.GroupVersionResource{Group: "operator.openshift.io", Version: "v1", Resource: "ingresscontrollers"}

var ingresscontrollersKind = schema.GroupVersionKind{Group: "operator.openshift.io", Version: "v1", Kind: "IngressController"}

// Get takes name of the ingressController, and returns the corresponding ingressController object, and an error if there is any.
func (c *FakeIngressControllers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingresscontrollersResource, c.ns, name), &operatorv1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.IngressController), err
}

// List takes label and field selectors, and returns the list of IngressControllers that match those selectors.
func (c *FakeIngressControllers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.IngressControllerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingresscontrollersResource, ingresscontrollersKind, c.ns, opts), &operatorv1.IngressControllerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.IngressControllerList{ListMeta: obj.(*operatorv1.IngressControllerList).ListMeta}
	for _, item := range obj.(*operatorv1.IngressControllerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressControllers.
func (c *FakeIngressControllers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingresscontrollersResource, c.ns, opts))

}

// Create takes the representation of a ingressController and creates it.  Returns the server's representation of the ingressController, and an error, if there is any.
func (c *FakeIngressControllers) Create(ctx context.Context, ingressController *operatorv1.IngressController, opts v1.CreateOptions) (result *operatorv1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingresscontrollersResource, c.ns, ingressController), &operatorv1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.IngressController), err
}

// Update takes the representation of a ingressController and updates it. Returns the server's representation of the ingressController, and an error, if there is any.
func (c *FakeIngressControllers) Update(ctx context.Context, ingressController *operatorv1.IngressController, opts v1.UpdateOptions) (result *operatorv1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingresscontrollersResource, c.ns, ingressController), &operatorv1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.IngressController), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIngressControllers) UpdateStatus(ctx context.Context, ingressController *operatorv1.IngressController, opts v1.UpdateOptions) (*operatorv1.IngressController, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ingresscontrollersResource, "status", c.ns, ingressController), &operatorv1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.IngressController), err
}

// Delete takes name of the ingressController and deletes it. Returns an error if one occurs.
func (c *FakeIngressControllers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ingresscontrollersResource, c.ns, name, opts), &operatorv1.IngressController{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressControllers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingresscontrollersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.IngressControllerList{})
	return err
}

// Patch applies the patch and returns the patched ingressController.
func (c *FakeIngressControllers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.IngressController, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingresscontrollersResource, c.ns, name, pt, data, subresources...), &operatorv1.IngressController{})

	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.IngressController), err
}
