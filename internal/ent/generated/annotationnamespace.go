// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotationnamespace"
	"go.infratographer.com/x/gidx"
)

// Representation of an annotation namespace. Annotation namespaces are used group annotation data that is provided by the same source and uses the same schema.
type AnnotationNamespace struct {
	config `json:"-"`
	// ID of the ent.
	// The ID for the annotation namespace.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the annotation namespace.
	Name string `json:"name,omitempty"`
	// The ID for the owner for this annotation namespace.
	OwnerID gidx.PrefixedID `json:"owner_id,omitempty"`
	// Flag for if this namespace is private.
	Private bool `json:"private,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnnotationNamespaceQuery when eager-loading is set.
	Edges        AnnotationNamespaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AnnotationNamespaceEdges holds the relations/edges for other nodes in the graph.
type AnnotationNamespaceEdges struct {
	// Annotations holds the value of the annotations edge.
	Annotations []*Annotation `json:"annotations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedAnnotations map[string][]*Annotation
}

// AnnotationsOrErr returns the Annotations value or an error if the edge
// was not loaded in eager-loading.
func (e AnnotationNamespaceEdges) AnnotationsOrErr() ([]*Annotation, error) {
	if e.loadedTypes[0] {
		return e.Annotations, nil
	}
	return nil, &NotLoadedError{edge: "annotations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AnnotationNamespace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case annotationnamespace.FieldID, annotationnamespace.FieldOwnerID:
			values[i] = new(gidx.PrefixedID)
		case annotationnamespace.FieldPrivate:
			values[i] = new(sql.NullBool)
		case annotationnamespace.FieldName:
			values[i] = new(sql.NullString)
		case annotationnamespace.FieldCreatedAt, annotationnamespace.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AnnotationNamespace fields.
func (an *AnnotationNamespace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case annotationnamespace.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				an.ID = *value
			}
		case annotationnamespace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				an.CreatedAt = value.Time
			}
		case annotationnamespace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				an.UpdatedAt = value.Time
			}
		case annotationnamespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				an.Name = value.String
			}
		case annotationnamespace.FieldOwnerID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value != nil {
				an.OwnerID = *value
			}
		case annotationnamespace.FieldPrivate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field private", values[i])
			} else if value.Valid {
				an.Private = value.Bool
			}
		default:
			an.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AnnotationNamespace.
// This includes values selected through modifiers, order, etc.
func (an *AnnotationNamespace) Value(name string) (ent.Value, error) {
	return an.selectValues.Get(name)
}

// QueryAnnotations queries the "annotations" edge of the AnnotationNamespace entity.
func (an *AnnotationNamespace) QueryAnnotations() *AnnotationQuery {
	return NewAnnotationNamespaceClient(an.config).QueryAnnotations(an)
}

// Update returns a builder for updating this AnnotationNamespace.
// Note that you need to call AnnotationNamespace.Unwrap() before calling this method if this AnnotationNamespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (an *AnnotationNamespace) Update() *AnnotationNamespaceUpdateOne {
	return NewAnnotationNamespaceClient(an.config).UpdateOne(an)
}

// Unwrap unwraps the AnnotationNamespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (an *AnnotationNamespace) Unwrap() *AnnotationNamespace {
	_tx, ok := an.config.driver.(*txDriver)
	if !ok {
		panic("generated: AnnotationNamespace is not a transactional entity")
	}
	an.config.driver = _tx.drv
	return an
}

// String implements the fmt.Stringer.
func (an *AnnotationNamespace) String() string {
	var builder strings.Builder
	builder.WriteString("AnnotationNamespace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", an.ID))
	builder.WriteString("created_at=")
	builder.WriteString(an.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(an.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(an.Name)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", an.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("private=")
	builder.WriteString(fmt.Sprintf("%v", an.Private))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (an AnnotationNamespace) IsEntity() {}

// NamedAnnotations returns the Annotations named value or an error if the edge was not
// loaded in eager-loading with this name.
func (an *AnnotationNamespace) NamedAnnotations(name string) ([]*Annotation, error) {
	if an.Edges.namedAnnotations == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := an.Edges.namedAnnotations[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (an *AnnotationNamespace) appendNamedAnnotations(name string, edges ...*Annotation) {
	if an.Edges.namedAnnotations == nil {
		an.Edges.namedAnnotations = make(map[string][]*Annotation)
	}
	if len(edges) == 0 {
		an.Edges.namedAnnotations[name] = []*Annotation{}
	} else {
		an.Edges.namedAnnotations[name] = append(an.Edges.namedAnnotations[name], edges...)
	}
}

// AnnotationNamespaces is a parsable slice of AnnotationNamespace.
type AnnotationNamespaces []*AnnotationNamespace
