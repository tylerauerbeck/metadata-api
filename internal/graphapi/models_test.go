package graphapi_test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/metadata-api/internal/ent/generated"
)

type AnnotationNamespaceBuilder struct {
	Name     string
	TenantID gidx.PrefixedID
	Private  bool
}

func (b AnnotationNamespaceBuilder) MustNew(ctx context.Context) *ent.AnnotationNamespace {
	if b.Name == "" {
		b.Name = fmt.Sprintf("%s/%s", gofakeit.DomainName(), gofakeit.Fruit())
	}

	if b.TenantID == "" {
		b.TenantID = gidx.MustNewID(tenantPrefix)
	}

	return EntClient.AnnotationNamespace.Create().SetName(b.Name).SetTenantID(b.TenantID).SetPrivate(b.Private).SaveX(ctx)
}

type MetadataBuilder struct {
	NodeID gidx.PrefixedID
}

func (b MetadataBuilder) MustNew(ctx context.Context) *ent.Metadata {
	if b.NodeID == "" {
		b.NodeID = gidx.MustNewID("tstnode")
	}

	return EntClient.Metadata.Create().SetNodeID(b.NodeID).SaveX(ctx)
}

type AnnotationBuilder struct {
	Metadata            *ent.Metadata
	AnnotationNamespace *ent.AnnotationNamespace
	Data                json.RawMessage
}

func (b AnnotationBuilder) MustNew(ctx context.Context) *ent.Annotation {
	if b.Metadata == nil {
		b.Metadata = MetadataBuilder{}.MustNew(ctx)
	}

	if b.AnnotationNamespace == nil {
		b.AnnotationNamespace = AnnotationNamespaceBuilder{}.MustNew(ctx)
	}

	if b.Data == nil {
		jsonData, err := gofakeit.JSON(nil)
		errPanic("generating random json", err)

		b.Data = json.RawMessage(jsonData)
	}

	return EntClient.Annotation.Create().SetMetadata(b.Metadata).SetNamespace(b.AnnotationNamespace).SetData(b.Data).SaveX(ctx)
}
