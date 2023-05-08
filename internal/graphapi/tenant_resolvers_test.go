package graphapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/metadata-api/internal/ent/generated"
	"go.infratographer.com/metadata-api/internal/graphclient"
)

func TestTenantAnnotationNamespaces(t *testing.T) {
	ctx := context.Background()
	tenantID := gidx.MustNewID(tenantPrefix)
	antColors := AnnotationNamespaceBuilder{TenantID: tenantID, Name: "nicole.dev/colors"}.MustNew(ctx)
	antPeople := AnnotationNamespaceBuilder{TenantID: tenantID, Name: "nicole.dev/people"}.MustNew(ctx)
	antPlaces := AnnotationNamespaceBuilder{TenantID: tenantID, Name: "nicole.dev/places"}.MustNew(ctx)
	AnnotationNamespaceBuilder{}.MustNew(ctx)
	// Update antColors so it's updated at is most recent
	antColors.Update().SaveX(ctx)

	testCases := []struct {
		TestName      string
		OrderBy       *graphclient.AnnotationNamespaceOrder
		TenantID      gidx.PrefixedID
		ResponseOrder []*ent.AnnotationNamespace
		errorMsg      string
	}{
		{
			TestName:      "Ordered By NAME ASC",
			OrderBy:       &graphclient.AnnotationNamespaceOrder{Field: "NAME", Direction: "ASC"},
			TenantID:      tenantID,
			ResponseOrder: []*ent.AnnotationNamespace{antColors, antPeople, antPlaces},
		},
		{
			TestName:      "Ordered By NAME DESC",
			OrderBy:       &graphclient.AnnotationNamespaceOrder{Field: "NAME", Direction: "DESC"},
			TenantID:      tenantID,
			ResponseOrder: []*ent.AnnotationNamespace{antPlaces, antPeople, antColors},
		},
		{
			TestName:      "Ordered By CREATED_AT ASC",
			OrderBy:       &graphclient.AnnotationNamespaceOrder{Field: "CREATED_AT", Direction: "ASC"},
			TenantID:      tenantID,
			ResponseOrder: []*ent.AnnotationNamespace{antColors, antPeople, antPlaces},
		},
		{
			TestName:      "Ordered By CREATED_AT DESC",
			OrderBy:       &graphclient.AnnotationNamespaceOrder{Field: "CREATED_AT", Direction: "DESC"},
			TenantID:      tenantID,
			ResponseOrder: []*ent.AnnotationNamespace{antPlaces, antPeople, antColors},
		},
		{
			TestName:      "Ordered By UPDATED_AT ASC",
			OrderBy:       &graphclient.AnnotationNamespaceOrder{Field: "UPDATED_AT", Direction: "ASC"},
			TenantID:      tenantID,
			ResponseOrder: []*ent.AnnotationNamespace{antPeople, antPlaces, antColors},
		},
		{
			TestName:      "Ordered By UPDATED_AT DESC",
			OrderBy:       &graphclient.AnnotationNamespaceOrder{Field: "UPDATED_AT", Direction: "DESC"},
			TenantID:      tenantID,
			ResponseOrder: []*ent.AnnotationNamespace{antColors, antPlaces, antPeople},
		},
		{
			TestName:      "No Annotation Namespaces for Tenant",
			TenantID:      gidx.MustNewID(tenantPrefix),
			ResponseOrder: []*ent.AnnotationNamespace{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient().GetTenantAnnotationNamespaces(ctx, tt.TenantID, tt.OrderBy)

			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)

				return
			}

			require.Len(t, resp.Entities[0].AnnotationNamespaces.Edges, len(tt.ResponseOrder))
			for i, lb := range tt.ResponseOrder {
				respNS := resp.Entities[0].AnnotationNamespaces.Edges[i].Node
				assert.Equal(t, lb.ID, respNS.ID)
				assert.Equal(t, lb.Name, respNS.Name)
			}
		})
	}
}
