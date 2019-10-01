// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/api/console/v1"
	scheme "github.com/openshift/client-go/console/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConsoleNotificationsGetter has a method to return a ConsoleNotificationInterface.
// A group's client should implement this interface.
type ConsoleNotificationsGetter interface {
	ConsoleNotifications() ConsoleNotificationInterface
}

// ConsoleNotificationInterface has methods to work with ConsoleNotification resources.
type ConsoleNotificationInterface interface {
	Create(*v1.ConsoleNotification) (*v1.ConsoleNotification, error)
	Update(*v1.ConsoleNotification) (*v1.ConsoleNotification, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ConsoleNotification, error)
	List(opts metav1.ListOptions) (*v1.ConsoleNotificationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ConsoleNotification, err error)
	ConsoleNotificationExpansion
}

// consoleNotifications implements ConsoleNotificationInterface
type consoleNotifications struct {
	client rest.Interface
}

// newConsoleNotifications returns a ConsoleNotifications
func newConsoleNotifications(c *ConsoleV1Client) *consoleNotifications {
	return &consoleNotifications{
		client: c.RESTClient(),
	}
}

// Get takes name of the consoleNotification, and returns the corresponding consoleNotification object, and an error if there is any.
func (c *consoleNotifications) Get(name string, options metav1.GetOptions) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Get().
		Resource("consolenotifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConsoleNotifications that match those selectors.
func (c *consoleNotifications) List(opts metav1.ListOptions) (result *v1.ConsoleNotificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ConsoleNotificationList{}
	err = c.client.Get().
		Resource("consolenotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested consoleNotifications.
func (c *consoleNotifications) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("consolenotifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a consoleNotification and creates it.  Returns the server's representation of the consoleNotification, and an error, if there is any.
func (c *consoleNotifications) Create(consoleNotification *v1.ConsoleNotification) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Post().
		Resource("consolenotifications").
		Body(consoleNotification).
		Do().
		Into(result)
	return
}

// Update takes the representation of a consoleNotification and updates it. Returns the server's representation of the consoleNotification, and an error, if there is any.
func (c *consoleNotifications) Update(consoleNotification *v1.ConsoleNotification) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Put().
		Resource("consolenotifications").
		Name(consoleNotification.Name).
		Body(consoleNotification).
		Do().
		Into(result)
	return
}

// Delete takes name of the consoleNotification and deletes it. Returns an error if one occurs.
func (c *consoleNotifications) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("consolenotifications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *consoleNotifications) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("consolenotifications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched consoleNotification.
func (c *consoleNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ConsoleNotification, err error) {
	result = &v1.ConsoleNotification{}
	err = c.client.Patch(pt).
		Resource("consolenotifications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
