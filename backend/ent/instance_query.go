// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marcustut/fyp/backend/ent/instance"
	"github.com/marcustut/fyp/backend/ent/predicate"
	"github.com/marcustut/fyp/backend/ent/schema/ulid"
	"github.com/marcustut/fyp/backend/ent/slide"
	"github.com/marcustut/fyp/backend/ent/user"
)

// InstanceQuery is the builder for querying Instance entities.
type InstanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Instance
	// eager-loading edges.
	withUser  *UserQuery
	withSlide *SlideQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InstanceQuery builder.
func (iq *InstanceQuery) Where(ps ...predicate.Instance) *InstanceQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InstanceQuery) Limit(limit int) *InstanceQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InstanceQuery) Offset(offset int) *InstanceQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InstanceQuery) Unique(unique bool) *InstanceQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *InstanceQuery) Order(o ...OrderFunc) *InstanceQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryUser chains the current query on the "user" edge.
func (iq *InstanceQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instance.UserTable, instance.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySlide chains the current query on the "slide" edge.
func (iq *InstanceQuery) QuerySlide() *SlideQuery {
	query := &SlideQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, selector),
			sqlgraph.To(slide.Table, slide.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, instance.SlideTable, instance.SlideColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Instance entity from the query.
// Returns a *NotFoundError when no Instance was found.
func (iq *InstanceQuery) First(ctx context.Context) (*Instance, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{instance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InstanceQuery) FirstX(ctx context.Context) *Instance {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Instance ID from the query.
// Returns a *NotFoundError when no Instance ID was found.
func (iq *InstanceQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{instance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InstanceQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Instance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Instance entity is not found.
// Returns a *NotFoundError when no Instance entities are found.
func (iq *InstanceQuery) Only(ctx context.Context) (*Instance, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{instance.Label}
	default:
		return nil, &NotSingularError{instance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InstanceQuery) OnlyX(ctx context.Context) *Instance {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Instance ID in the query.
// Returns a *NotSingularError when exactly one Instance ID is not found.
// Returns a *NotFoundError when no entities are found.
func (iq *InstanceQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = &NotSingularError{instance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InstanceQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Instances.
func (iq *InstanceQuery) All(ctx context.Context) ([]*Instance, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InstanceQuery) AllX(ctx context.Context) []*Instance {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Instance IDs.
func (iq *InstanceQuery) IDs(ctx context.Context) ([]ulid.ID, error) {
	var ids []ulid.ID
	if err := iq.Select(instance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InstanceQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InstanceQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InstanceQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InstanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InstanceQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InstanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InstanceQuery) Clone() *InstanceQuery {
	if iq == nil {
		return nil
	}
	return &InstanceQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]OrderFunc{}, iq.order...),
		predicates: append([]predicate.Instance{}, iq.predicates...),
		withUser:   iq.withUser.Clone(),
		withSlide:  iq.withSlide.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstanceQuery) WithUser(opts ...func(*UserQuery)) *InstanceQuery {
	query := &UserQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// WithSlide tells the query-builder to eager-load the nodes that are connected to
// the "slide" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InstanceQuery) WithSlide(opts ...func(*SlideQuery)) *InstanceQuery {
	query := &SlideQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withSlide = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InstanceID string `json:"instance_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Instance.Query().
//		GroupBy(instance.FieldInstanceID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *InstanceQuery) GroupBy(field string, fields ...string) *InstanceGroupBy {
	group := &InstanceGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InstanceID string `json:"instance_id,omitempty"`
//	}
//
//	client.Instance.Query().
//		Select(instance.FieldInstanceID).
//		Scan(ctx, &v)
//
func (iq *InstanceQuery) Select(fields ...string) *InstanceSelect {
	iq.fields = append(iq.fields, fields...)
	return &InstanceSelect{InstanceQuery: iq}
}

func (iq *InstanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !instance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InstanceQuery) sqlAll(ctx context.Context) ([]*Instance, error) {
	var (
		nodes       = []*Instance{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withUser != nil,
			iq.withSlide != nil,
		}
	)
	if iq.withUser != nil || iq.withSlide != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, instance.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Instance{config: iq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withUser; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*Instance)
		for i := range nodes {
			if nodes[i].user_instances == nil {
				continue
			}
			fk := *nodes[i].user_instances
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_instances" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := iq.withSlide; query != nil {
		ids := make([]ulid.ID, 0, len(nodes))
		nodeids := make(map[ulid.ID][]*Instance)
		for i := range nodes {
			if nodes[i].slide_instance == nil {
				continue
			}
			fk := *nodes[i].slide_instance
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(slide.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "slide_instance" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Slide = n
			}
		}
	}

	return nodes, nil
}

func (iq *InstanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InstanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iq *InstanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instance.Table,
			Columns: instance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: instance.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instance.FieldID)
		for i := range fields {
			if fields[i] != instance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InstanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(instance.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = instance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InstanceGroupBy is the group-by builder for Instance entities.
type InstanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InstanceGroupBy) Aggregate(fns ...AggregateFunc) *InstanceGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *InstanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *InstanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *InstanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = igb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (igb *InstanceGroupBy) StringX(ctx context.Context) string {
	v, err := igb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *InstanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = igb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (igb *InstanceGroupBy) IntX(ctx context.Context) int {
	v, err := igb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *InstanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = igb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (igb *InstanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := igb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InstanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *InstanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (igb *InstanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = igb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (igb *InstanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := igb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *InstanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !instance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InstanceGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// InstanceSelect is the builder for selecting fields of Instance entities.
type InstanceSelect struct {
	*InstanceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *InstanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.InstanceQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *InstanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *InstanceSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = is.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (is *InstanceSelect) StringX(ctx context.Context) string {
	v, err := is.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *InstanceSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = is.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (is *InstanceSelect) IntX(ctx context.Context) int {
	v, err := is.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *InstanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = is.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (is *InstanceSelect) Float64X(ctx context.Context) float64 {
	v, err := is.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InstanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *InstanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (is *InstanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = is.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{instance.Label}
	default:
		err = fmt.Errorf("ent: InstanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (is *InstanceSelect) BoolX(ctx context.Context) bool {
	v, err := is.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *InstanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
