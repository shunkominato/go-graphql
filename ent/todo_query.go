// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"story.com/story/app/ent/predicate"
	"story.com/story/app/ent/todo"
	"story.com/story/app/ent/todostatus"
	"story.com/story/app/ent/user"
)

// TodoQuery is the builder for querying Todo entities.
type TodoQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.Todo
	withTodoStatu *TodoStatusQuery
	withUser      *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TodoQuery builder.
func (tq *TodoQuery) Where(ps ...predicate.Todo) *TodoQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TodoQuery) Limit(limit int) *TodoQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TodoQuery) Offset(offset int) *TodoQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TodoQuery) Unique(unique bool) *TodoQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TodoQuery) Order(o ...OrderFunc) *TodoQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryTodoStatu chains the current query on the "todo_statu" edge.
func (tq *TodoQuery) QueryTodoStatu() *TodoStatusQuery {
	query := (&TodoStatusClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, selector),
			sqlgraph.To(todostatus.Table, todostatus.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.TodoStatuTable, todo.TodoStatuColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (tq *TodoQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.UserTable, todo.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Todo entity from the query.
// Returns a *NotFoundError when no Todo was found.
func (tq *TodoQuery) First(ctx context.Context) (*Todo, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{todo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TodoQuery) FirstX(ctx context.Context) *Todo {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Todo ID from the query.
// Returns a *NotFoundError when no Todo ID was found.
func (tq *TodoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{todo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TodoQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Todo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Todo entity is found.
// Returns a *NotFoundError when no Todo entities are found.
func (tq *TodoQuery) Only(ctx context.Context) (*Todo, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{todo.Label}
	default:
		return nil, &NotSingularError{todo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TodoQuery) OnlyX(ctx context.Context) *Todo {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Todo ID in the query.
// Returns a *NotSingularError when more than one Todo ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TodoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{todo.Label}
	default:
		err = &NotSingularError{todo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TodoQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Todos.
func (tq *TodoQuery) All(ctx context.Context) ([]*Todo, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Todo, *TodoQuery]()
	return withInterceptors[[]*Todo](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TodoQuery) AllX(ctx context.Context) []*Todo {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Todo IDs.
func (tq *TodoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(todo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TodoQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TodoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TodoQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TodoQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TodoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TodoQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TodoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TodoQuery) Clone() *TodoQuery {
	if tq == nil {
		return nil
	}
	return &TodoQuery{
		config:        tq.config,
		ctx:           tq.ctx.Clone(),
		order:         append([]OrderFunc{}, tq.order...),
		inters:        append([]Interceptor{}, tq.inters...),
		predicates:    append([]predicate.Todo{}, tq.predicates...),
		withTodoStatu: tq.withTodoStatu.Clone(),
		withUser:      tq.withUser.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithTodoStatu tells the query-builder to eager-load the nodes that are connected to
// the "todo_statu" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TodoQuery) WithTodoStatu(opts ...func(*TodoStatusQuery)) *TodoQuery {
	query := (&TodoStatusClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withTodoStatu = query
	return tq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TodoQuery) WithUser(opts ...func(*UserQuery)) *TodoQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Todo string `json:"todo,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Todo.Query().
//		GroupBy(todo.FieldTodo).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TodoQuery) GroupBy(field string, fields ...string) *TodoGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TodoGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = todo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Todo string `json:"todo,omitempty"`
//	}
//
//	client.Todo.Query().
//		Select(todo.FieldTodo).
//		Scan(ctx, &v)
func (tq *TodoQuery) Select(fields ...string) *TodoSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TodoSelect{TodoQuery: tq}
	sbuild.label = todo.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TodoSelect configured with the given aggregations.
func (tq *TodoQuery) Aggregate(fns ...AggregateFunc) *TodoSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TodoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !todo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TodoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Todo, error) {
	var (
		nodes       = []*Todo{}
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withTodoStatu != nil,
			tq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Todo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Todo{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withTodoStatu; query != nil {
		if err := tq.loadTodoStatu(ctx, query, nodes, nil,
			func(n *Todo, e *TodoStatus) { n.Edges.TodoStatu = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withUser; query != nil {
		if err := tq.loadUser(ctx, query, nodes, nil,
			func(n *Todo, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TodoQuery) loadTodoStatu(ctx context.Context, query *TodoStatusQuery, nodes []*Todo, init func(*Todo), assign func(*Todo, *TodoStatus)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Todo)
	for i := range nodes {
		fk := nodes[i].TodoStatusesID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(todostatus.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "todo_statuses_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TodoQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Todo, init func(*Todo), assign func(*Todo, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Todo)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TodoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TodoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for i := range fields {
			if fields[i] != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TodoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(todo.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = todo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TodoGroupBy is the group-by builder for Todo entities.
type TodoGroupBy struct {
	selector
	build *TodoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TodoGroupBy) Aggregate(fns ...AggregateFunc) *TodoGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TodoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodoQuery, *TodoGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TodoGroupBy) sqlScan(ctx context.Context, root *TodoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TodoSelect is the builder for selecting fields of Todo entities.
type TodoSelect struct {
	*TodoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TodoSelect) Aggregate(fns ...AggregateFunc) *TodoSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TodoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodoQuery, *TodoSelect](ctx, ts.TodoQuery, ts, ts.inters, v)
}

func (ts *TodoSelect) sqlScan(ctx context.Context, root *TodoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
