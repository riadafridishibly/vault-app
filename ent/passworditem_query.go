// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/riadafridishibly/vault-app/ent/passworditem"
	"github.com/riadafridishibly/vault-app/ent/predicate"
)

// PasswordItemQuery is the builder for querying PasswordItem entities.
type PasswordItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PasswordItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PasswordItemQuery builder.
func (piq *PasswordItemQuery) Where(ps ...predicate.PasswordItem) *PasswordItemQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit adds a limit step to the query.
func (piq *PasswordItemQuery) Limit(limit int) *PasswordItemQuery {
	piq.limit = &limit
	return piq
}

// Offset adds an offset step to the query.
func (piq *PasswordItemQuery) Offset(offset int) *PasswordItemQuery {
	piq.offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *PasswordItemQuery) Unique(unique bool) *PasswordItemQuery {
	piq.unique = &unique
	return piq
}

// Order adds an order step to the query.
func (piq *PasswordItemQuery) Order(o ...OrderFunc) *PasswordItemQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// First returns the first PasswordItem entity from the query.
// Returns a *NotFoundError when no PasswordItem was found.
func (piq *PasswordItemQuery) First(ctx context.Context) (*PasswordItem, error) {
	nodes, err := piq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{passworditem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *PasswordItemQuery) FirstX(ctx context.Context) *PasswordItem {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PasswordItem ID from the query.
// Returns a *NotFoundError when no PasswordItem ID was found.
func (piq *PasswordItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{passworditem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *PasswordItemQuery) FirstIDX(ctx context.Context) int {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PasswordItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PasswordItem entity is found.
// Returns a *NotFoundError when no PasswordItem entities are found.
func (piq *PasswordItemQuery) Only(ctx context.Context) (*PasswordItem, error) {
	nodes, err := piq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{passworditem.Label}
	default:
		return nil, &NotSingularError{passworditem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *PasswordItemQuery) OnlyX(ctx context.Context) *PasswordItem {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PasswordItem ID in the query.
// Returns a *NotSingularError when more than one PasswordItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *PasswordItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{passworditem.Label}
	default:
		err = &NotSingularError{passworditem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *PasswordItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PasswordItems.
func (piq *PasswordItemQuery) All(ctx context.Context) ([]*PasswordItem, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return piq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (piq *PasswordItemQuery) AllX(ctx context.Context) []*PasswordItem {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PasswordItem IDs.
func (piq *PasswordItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := piq.Select(passworditem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *PasswordItemQuery) IDsX(ctx context.Context) []int {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *PasswordItemQuery) Count(ctx context.Context) (int, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return piq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (piq *PasswordItemQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *PasswordItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return piq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *PasswordItemQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PasswordItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *PasswordItemQuery) Clone() *PasswordItemQuery {
	if piq == nil {
		return nil
	}
	return &PasswordItemQuery{
		config:     piq.config,
		limit:      piq.limit,
		offset:     piq.offset,
		order:      append([]OrderFunc{}, piq.order...),
		predicates: append([]predicate.PasswordItem{}, piq.predicates...),
		// clone intermediate query.
		sql:    piq.sql.Clone(),
		path:   piq.path,
		unique: piq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PasswordItem.Query().
//		GroupBy(passworditem.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *PasswordItemQuery) GroupBy(field string, fields ...string) *PasswordItemGroupBy {
	grbuild := &PasswordItemGroupBy{config: piq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return piq.sqlQuery(ctx), nil
	}
	grbuild.label = passworditem.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.PasswordItem.Query().
//		Select(passworditem.FieldCreateTime).
//		Scan(ctx, &v)
func (piq *PasswordItemQuery) Select(fields ...string) *PasswordItemSelect {
	piq.fields = append(piq.fields, fields...)
	selbuild := &PasswordItemSelect{PasswordItemQuery: piq}
	selbuild.label = passworditem.Label
	selbuild.flds, selbuild.scan = &piq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a PasswordItemSelect configured with the given aggregations.
func (piq *PasswordItemQuery) Aggregate(fns ...AggregateFunc) *PasswordItemSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *PasswordItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range piq.fields {
		if !passworditem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *PasswordItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PasswordItem, error) {
	var (
		nodes = []*PasswordItem{}
		_spec = piq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PasswordItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PasswordItem{config: piq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (piq *PasswordItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.fields
	if len(piq.fields) > 0 {
		_spec.Unique = piq.unique != nil && *piq.unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *PasswordItemQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (piq *PasswordItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passworditem.Table,
			Columns: passworditem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passworditem.FieldID,
			},
		},
		From:   piq.sql,
		Unique: true,
	}
	if unique := piq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := piq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passworditem.FieldID)
		for i := range fields {
			if fields[i] != passworditem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *PasswordItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(passworditem.Table)
	columns := piq.fields
	if len(columns) == 0 {
		columns = passworditem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.unique != nil && *piq.unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PasswordItemGroupBy is the group-by builder for PasswordItem entities.
type PasswordItemGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *PasswordItemGroupBy) Aggregate(fns ...AggregateFunc) *PasswordItemGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the group-by query and scans the result into the given value.
func (pigb *PasswordItemGroupBy) Scan(ctx context.Context, v any) error {
	query, err := pigb.path(ctx)
	if err != nil {
		return err
	}
	pigb.sql = query
	return pigb.sqlScan(ctx, v)
}

func (pigb *PasswordItemGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range pigb.fields {
		if !passworditem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pigb *PasswordItemGroupBy) sqlQuery() *sql.Selector {
	selector := pigb.sql.Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pigb.fields)+len(pigb.fns))
		for _, f := range pigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pigb.fields...)...)
}

// PasswordItemSelect is the builder for selecting fields of PasswordItem entities.
type PasswordItemSelect struct {
	*PasswordItemQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *PasswordItemSelect) Aggregate(fns ...AggregateFunc) *PasswordItemSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *PasswordItemSelect) Scan(ctx context.Context, v any) error {
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	pis.sql = pis.PasswordItemQuery.sqlQuery(ctx)
	return pis.sqlScan(ctx, v)
}

func (pis *PasswordItemSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(pis.sql))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		pis.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		pis.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := pis.sql.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
