// Code generated by solo-kit. DO NOT EDIT.

package v1alpha3

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionServiceEntryFunc func(original, desired *ServiceEntry) (bool, error)

type ServiceEntryReconciler interface {
	Reconcile(namespace string, desiredResources ServiceEntryList, transition TransitionServiceEntryFunc, opts clients.ListOpts) error
}

func serviceEntrysToResources(list ServiceEntryList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, serviceEntry := range list {
		resourceList = append(resourceList, serviceEntry)
	}
	return resourceList
}

func NewServiceEntryReconciler(client ServiceEntryClient) ServiceEntryReconciler {
	return &serviceEntryReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type serviceEntryReconciler struct {
	base reconcile.Reconciler
}

func (r *serviceEntryReconciler) Reconcile(namespace string, desiredResources ServiceEntryList, transition TransitionServiceEntryFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "serviceEntry_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*ServiceEntry), desired.(*ServiceEntry))
		}
	}
	return r.base.Reconcile(namespace, serviceEntrysToResources(desiredResources), transitionResources, opts)
}
