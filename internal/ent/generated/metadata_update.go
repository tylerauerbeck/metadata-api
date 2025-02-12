// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/metadata-api/internal/ent/generated/annotation"
	"go.infratographer.com/metadata-api/internal/ent/generated/metadata"
	"go.infratographer.com/metadata-api/internal/ent/generated/predicate"
	"go.infratographer.com/metadata-api/internal/ent/generated/status"
	"go.infratographer.com/x/gidx"
)

// MetadataUpdate is the builder for updating Metadata entities.
type MetadataUpdate struct {
	config
	hooks    []Hook
	mutation *MetadataMutation
}

// Where appends a list predicates to the MetadataUpdate builder.
func (mu *MetadataUpdate) Where(ps ...predicate.Metadata) *MetadataUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (mu *MetadataUpdate) AddAnnotationIDs(ids ...gidx.PrefixedID) *MetadataUpdate {
	mu.mutation.AddAnnotationIDs(ids...)
	return mu
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (mu *MetadataUpdate) AddAnnotations(a ...*Annotation) *MetadataUpdate {
	ids := make([]gidx.PrefixedID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mu.AddAnnotationIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (mu *MetadataUpdate) AddStatusIDs(ids ...gidx.PrefixedID) *MetadataUpdate {
	mu.mutation.AddStatusIDs(ids...)
	return mu
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (mu *MetadataUpdate) AddStatuses(s ...*Status) *MetadataUpdate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.AddStatusIDs(ids...)
}

// Mutation returns the MetadataMutation object of the builder.
func (mu *MetadataUpdate) Mutation() *MetadataMutation {
	return mu.mutation
}

// ClearAnnotations clears all "annotations" edges to the Annotation entity.
func (mu *MetadataUpdate) ClearAnnotations() *MetadataUpdate {
	mu.mutation.ClearAnnotations()
	return mu
}

// RemoveAnnotationIDs removes the "annotations" edge to Annotation entities by IDs.
func (mu *MetadataUpdate) RemoveAnnotationIDs(ids ...gidx.PrefixedID) *MetadataUpdate {
	mu.mutation.RemoveAnnotationIDs(ids...)
	return mu
}

// RemoveAnnotations removes "annotations" edges to Annotation entities.
func (mu *MetadataUpdate) RemoveAnnotations(a ...*Annotation) *MetadataUpdate {
	ids := make([]gidx.PrefixedID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return mu.RemoveAnnotationIDs(ids...)
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (mu *MetadataUpdate) ClearStatuses() *MetadataUpdate {
	mu.mutation.ClearStatuses()
	return mu
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (mu *MetadataUpdate) RemoveStatusIDs(ids ...gidx.PrefixedID) *MetadataUpdate {
	mu.mutation.RemoveStatusIDs(ids...)
	return mu
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (mu *MetadataUpdate) RemoveStatuses(s ...*Status) *MetadataUpdate {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.RemoveStatusIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MetadataUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MetadataUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MetadataUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MetadataUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MetadataUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := metadata.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

func (mu *MetadataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(metadata.FieldUpdatedAt, field.TypeTime, value)
	}
	if mu.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedAnnotationsIDs(); len(nodes) > 0 && !mu.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !mu.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MetadataUpdateOne is the builder for updating a single Metadata entity.
type MetadataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MetadataMutation
}

// AddAnnotationIDs adds the "annotations" edge to the Annotation entity by IDs.
func (muo *MetadataUpdateOne) AddAnnotationIDs(ids ...gidx.PrefixedID) *MetadataUpdateOne {
	muo.mutation.AddAnnotationIDs(ids...)
	return muo
}

// AddAnnotations adds the "annotations" edges to the Annotation entity.
func (muo *MetadataUpdateOne) AddAnnotations(a ...*Annotation) *MetadataUpdateOne {
	ids := make([]gidx.PrefixedID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return muo.AddAnnotationIDs(ids...)
}

// AddStatusIDs adds the "statuses" edge to the Status entity by IDs.
func (muo *MetadataUpdateOne) AddStatusIDs(ids ...gidx.PrefixedID) *MetadataUpdateOne {
	muo.mutation.AddStatusIDs(ids...)
	return muo
}

// AddStatuses adds the "statuses" edges to the Status entity.
func (muo *MetadataUpdateOne) AddStatuses(s ...*Status) *MetadataUpdateOne {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.AddStatusIDs(ids...)
}

// Mutation returns the MetadataMutation object of the builder.
func (muo *MetadataUpdateOne) Mutation() *MetadataMutation {
	return muo.mutation
}

// ClearAnnotations clears all "annotations" edges to the Annotation entity.
func (muo *MetadataUpdateOne) ClearAnnotations() *MetadataUpdateOne {
	muo.mutation.ClearAnnotations()
	return muo
}

// RemoveAnnotationIDs removes the "annotations" edge to Annotation entities by IDs.
func (muo *MetadataUpdateOne) RemoveAnnotationIDs(ids ...gidx.PrefixedID) *MetadataUpdateOne {
	muo.mutation.RemoveAnnotationIDs(ids...)
	return muo
}

// RemoveAnnotations removes "annotations" edges to Annotation entities.
func (muo *MetadataUpdateOne) RemoveAnnotations(a ...*Annotation) *MetadataUpdateOne {
	ids := make([]gidx.PrefixedID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return muo.RemoveAnnotationIDs(ids...)
}

// ClearStatuses clears all "statuses" edges to the Status entity.
func (muo *MetadataUpdateOne) ClearStatuses() *MetadataUpdateOne {
	muo.mutation.ClearStatuses()
	return muo
}

// RemoveStatusIDs removes the "statuses" edge to Status entities by IDs.
func (muo *MetadataUpdateOne) RemoveStatusIDs(ids ...gidx.PrefixedID) *MetadataUpdateOne {
	muo.mutation.RemoveStatusIDs(ids...)
	return muo
}

// RemoveStatuses removes "statuses" edges to Status entities.
func (muo *MetadataUpdateOne) RemoveStatuses(s ...*Status) *MetadataUpdateOne {
	ids := make([]gidx.PrefixedID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.RemoveStatusIDs(ids...)
}

// Where appends a list predicates to the MetadataUpdate builder.
func (muo *MetadataUpdateOne) Where(ps ...predicate.Metadata) *MetadataUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MetadataUpdateOne) Select(field string, fields ...string) *MetadataUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Metadata entity.
func (muo *MetadataUpdateOne) Save(ctx context.Context) (*Metadata, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MetadataUpdateOne) SaveX(ctx context.Context) *Metadata {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MetadataUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MetadataUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MetadataUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := metadata.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

func (muo *MetadataUpdateOne) sqlSave(ctx context.Context) (_node *Metadata, err error) {
	_spec := sqlgraph.NewUpdateSpec(metadata.Table, metadata.Columns, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Metadata.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, metadata.FieldID)
		for _, f := range fields {
			if !metadata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != metadata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(metadata.FieldUpdatedAt, field.TypeTime, value)
	}
	if muo.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedAnnotationsIDs(); len(nodes) > 0 && !muo.mutation.AnnotationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.AnnotationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.AnnotationsTable,
			Columns: []string{metadata.AnnotationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(annotation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedStatusesIDs(); len(nodes) > 0 && !muo.mutation.StatusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.StatusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.StatusesTable,
			Columns: []string{metadata.StatusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Metadata{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{metadata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
