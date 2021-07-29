/*
Copyright 2020 - 2021 Crunchy Data Solutions, Inc.
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

package fake

import (
	"context"

	crunchydatacomv1 "github.com/percona/percona-postgresql-operator/pkg/apis/crunchydata.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePgclusters implements PgclusterInterface
type FakePerconaPGClusters struct {
	Fake *FakeCrunchydataV1
	ns   string
}

var perconapgclustersResource = schema.GroupVersionResource{Group: "crunchydata.com", Version: "v1", Resource: "pgclusters"}

var perconapgclustersKind = schema.GroupVersionKind{Group: "crunchydata.com", Version: "v1", Kind: "Pgcluster"}

// Get takes name of the pgcluster, and returns the corresponding pgcluster object, and an error if there is any.
func (c *FakePerconaPGClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *crunchydatacomv1.PerconaPGCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pgclustersResource, c.ns, name), &crunchydatacomv1.Pgcluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crunchydatacomv1.PerconaPGCluster), err
}

// List takes label and field selectors, and returns the list of Pgclusters that match those selectors.
func (c *FakePerconaPGClusters) List(ctx context.Context, opts v1.ListOptions) (result *crunchydatacomv1.PerconaPGClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pgclustersResource, pgclustersKind, c.ns, opts), &crunchydatacomv1.PgclusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &crunchydatacomv1.PerconaPGClusterList{ListMeta: obj.(*crunchydatacomv1.PgclusterList).ListMeta}
	for _, item := range obj.(*crunchydatacomv1.PerconaPGClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pgclusters.
func (c *FakePerconaPGClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pgclustersResource, c.ns, opts))

}

// Create takes the representation of a pgcluster and creates it.  Returns the server's representation of the pgcluster, and an error, if there is any.
func (c *FakePerconaPGClusters) Create(ctx context.Context, pgcluster *crunchydatacomv1.PerconaPGCluster, opts v1.CreateOptions) (result *crunchydatacomv1.PerconaPGCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pgclustersResource, c.ns, pgcluster), &crunchydatacomv1.Pgcluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crunchydatacomv1.PerconaPGCluster), err
}

// Update takes the representation of a pgcluster and updates it. Returns the server's representation of the pgcluster, and an error, if there is any.
func (c *FakePerconaPGClusters) Update(ctx context.Context, pgcluster *crunchydatacomv1.PerconaPGCluster, opts v1.UpdateOptions) (result *crunchydatacomv1.PerconaPGCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pgclustersResource, c.ns, pgcluster), &crunchydatacomv1.Pgcluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crunchydatacomv1.PerconaPGCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePerconaPGClusters) UpdateStatus(ctx context.Context, pgcluster *crunchydatacomv1.PerconaPGCluster, opts v1.UpdateOptions) (*crunchydatacomv1.PerconaPGCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pgclustersResource, "status", c.ns, pgcluster), &crunchydatacomv1.Pgcluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crunchydatacomv1.PerconaPGCluster), err
}

// Delete takes name of the pgcluster and deletes it. Returns an error if one occurs.
func (c *FakePerconaPGClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pgclustersResource, c.ns, name), &crunchydatacomv1.PerconaPGCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePerconaPGClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pgclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &crunchydatacomv1.PerconaPGClusterList{})
	return err
}

// Patch applies the patch and returns the patched pgcluster.
func (c *FakePerconaPGClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *crunchydatacomv1.PerconaPGCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pgclustersResource, c.ns, name, pt, data, subresources...), &crunchydatacomv1.PerconaPGCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*crunchydatacomv1.PerconaPGCluster), err
}