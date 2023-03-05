// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"story.com/story/app/ent/predicate"
	"story.com/story/app/ent/todo"
	"story.com/story/app/ent/todostatus"
)

// TodoStatusQuery is the builder for querying TodoStatus entities.
type TodoStatusQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.TodoStatus
	withTodos  *TodoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TodoStatusQuery builder.
func (tsq *TodoStatusQuery) Where(ps ...predicate.TodoStatus) *TodoStatusQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit the number of records to be returned by this query.
func (tsq *TodoStatusQuery) Limit(limit int) *TodoStatusQuery {
	tsq.ctx.Limit = &limit
	return tsq
}

// Offset to start from.
func (tsq *TodoStatusQuery) Offset(offset int) *TodoStatusQuery {
	tsq.ctx.Offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *TodoStatusQuery) Unique(unique bool) *TodoStatusQuery {
	tsq.ctx.Unique = &unique
	return tsq
}

// Order specifies how the records should be ordered.
func (tsq *TodoStatusQuery) Order(o ...OrderFunc) *TodoStatusQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// QueryTodos chains the current query on the "todos" edge.
func (tsq *TodoStatusQuery) QueryTodos() *TodoQuery {
	query := (&TodoClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(todostatus.Table, todostatus.FieldID, selector),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, todostatus.TodosTable, todostatus.TodosColumn),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TodoStatus entity from the query.
// Returns a *NotFoundError when no TodoStatus was found.
func (tsq *TodoStatusQuery) First(ctx context.Context) (*TodoStatus, error) {
	nodes, err := tsq.Limit(1).All(setContextOp(ctx, tsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{todostatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *TodoStatusQuery) FirstX(ctx context.Context) *TodoStatus {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TodoStatus ID from the query.
// Returns a *NotFoundError when no TodoStatus ID was found.
func (tsq *TodoStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(1).IDs(setContextOp(ctx, tsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{todostatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *TodoStatusQuery) FirstIDX(ctx context.Context) int {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TodoStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TodoStatus entity is found.
// Returns a *NotFoundError when no TodoStatus entities are found.
func (tsq *TodoStatusQuery) Only(ctx context.Context) (*TodoStatus, error) {
	nodes, err := tsq.Limit(2).All(setContextOp(ctx, tsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{todostatus.Label}
	default:
		return nil, &NotSingularError{todostatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *TodoStatusQuery) OnlyX(ctx context.Context) *TodoStatus {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TodoStatus ID in the query.
// Returns a *NotSingularError when more than one TodoStatus ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *TodoStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(2).IDs(setContextOp(ctx, tsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{todostatus.Label}
	default:
		err = &NotSingularError{todostatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *TodoStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TodoStatusSlice.
func (tsq *TodoStatusQuery) All(ctx context.Context) ([]*TodoStatus, error) {
	ctx = setContextOp(ctx, tsq.ctx, "All")
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TodoStatus, *TodoStatusQuery]()
	return withInterceptors[[]*TodoStatus](ctx, tsq, qr, tsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsq *TodoStatusQuery) AllX(ctx context.Context) []*TodoStatus {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TodoStatus IDs.
func (tsq *TodoStatusQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tsq.ctx.Unique == nil && tsq.path != nil {
		tsq.Unique(true)
	}
	ctx = setContextOp(ctx, tsq.ctx, "IDs")
	if err = tsq.Select(todostatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *TodoStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *TodoStatusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Count")
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsq, querierCount[*TodoStatusQuery](), tsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *TodoStatusQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *TodoStatusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Exist")
	switch _, err := tsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *TodoStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TodoStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *TodoStatusQuery) Clone() *TodoStatusQuery {
	if tsq == nil {
		return nil
	}
	return &TodoStatusQuery{
		config:     tsq.config,
		ctx:        tsq.ctx.Clone(),
		order:      append([]OrderFunc{}, tsq.order...),
		inters:     append([]Interceptor{}, tsq.inters...),
		predicates: append([]predicate.TodoStatus{}, tsq.predicates...),
		withTodos:  tsq.withTodos.Clone(),
		// clone intermediate query.
		sql:  tsq.sql.Clone(),
		path: tsq.path,
	}
}

// WithTodos tells the query-builder to eager-load the nodes that are connected to
// the "todos" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TodoStatusQuery) WithTodos(opts ...func(*TodoQuery)) *TodoStatusQuery {
	query := (&TodoClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withTodos = query
	return tsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TodoStatus.Query().
//		GroupBy(todostatus.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tsq *TodoStatusQuery) GroupBy(field string, fields ...string) *TodoStatusGroupBy {
	tsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TodoStatusGroupBy{build: tsq}
	grbuild.flds = &tsq.ctx.Fields
	grbuild.label = todostatus.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//	}
//
//	client.TodoStatus.Query().
//		Select(todostatus.FieldStatus).
//		Scan(ctx, &v)
func (tsq *TodoStatusQuery) Select(fields ...string) *TodoStatusSelect {
	tsq.ctx.Fields = append(tsq.ctx.Fields, fields...)
	sbuild := &TodoStatusSelect{TodoStatusQuery: tsq}
	sbuild.label = todostatus.Label
	sbuild.flds, sbuild.scan = &tsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TodoStatusSelect configured with the given aggregations.
func (tsq *TodoStatusQuery) Aggregate(fns ...AggregateFunc) *TodoStatusSelect {
	return tsq.Select().Aggregate(fns...)
}

func (tsq *TodoStatusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsq.ctx.Fields {
		if !todostatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	return nil
}

func (tsq *TodoStatusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TodoStatus, error) {
	var (
		nodes       = []*TodoStatus{}
		_spec       = tsq.querySpec()
		loadedTypes = [1]bool{
			tsq.withTodos != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TodoStatus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TodoStatus{config: tsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tsq.withTodos; query != nil {
		if err := tsq.loadTodos(ctx, query, nodes,
			func(n *TodoStatus) { n.Edges.Todos = []*Todo{} },
			func(n *TodoStatus, e *Todo) { n.Edges.Todos = append(n.Edges.Todos, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tsq *TodoStatusQuery) loadTodos(ctx context.Context, query *TodoQuery, nodes []*TodoStatus, init func(*TodoStatus), assign func(*TodoStatus, *Todo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TodoStatus)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Todo(func(s *sql.Selector) {
		s.Where(sql.InValues(todostatus.TodosColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TodoStatusesID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "todo_statuses_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tsq *TodoStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Columns = tsq.ctx.Fields
	if len(tsq.ctx.Fields) > 0 {
		_spec.Unique = tsq.ctx.Unique != nil && *tsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *TodoStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(todostatus.Table, todostatus.Columns, sqlgraph.NewFieldSpec(todostatus.FieldID, field.TypeInt))
	_spec.From = tsq.sql
	if unique := tsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsq.path != nil {
		_spec.Unique = true
	}
	if fields := tsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todostatus.FieldID)
		for i := range fields {
			if fields[i] != todostatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *TodoStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(todostatus.Table)
	columns := tsq.ctx.Fields
	if len(columns) == 0 {
		columns = todostatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.ctx.Unique != nil && *tsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TodoStatusGroupBy is the group-by builder for TodoStatus entities.
type TodoStatusGroupBy struct {
	selector
	build *TodoStatusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *TodoStatusGroupBy) Aggregate(fns ...AggregateFunc) *TodoStatusGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsgb *TodoStatusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsgb.build.ctx, "GroupBy")
	if err := tsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodoStatusQuery, *TodoStatusGroupBy](ctx, tsgb.build, tsgb, tsgb.build.inters, v)
}

func (tsgb *TodoStatusGroupBy) sqlScan(ctx context.Context, root *TodoStatusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsgb.flds)+len(tsgb.fns))
		for _, f := range *tsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TodoStatusSelect is the builder for selecting fields of TodoStatus entities.
type TodoStatusSelect struct {
	*TodoStatusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tss *TodoStatusSelect) Aggregate(fns ...AggregateFunc) *TodoStatusSelect {
	tss.fns = append(tss.fns, fns...)
	return tss
}

// Scan applies the selector query and scans the result into the given value.
func (tss *TodoStatusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tss.ctx, "Select")
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodoStatusQuery, *TodoStatusSelect](ctx, tss.TodoStatusQuery, tss, tss.inters, v)
}

func (tss *TodoStatusSelect) sqlScan(ctx context.Context, root *TodoStatusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tss.fns))
	for _, fn := range tss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}