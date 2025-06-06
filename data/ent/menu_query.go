// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"coolfun-admin/data/ent/menu"
	"coolfun-admin/data/ent/menuparam"
	"coolfun-admin/data/ent/predicate"
	"coolfun-admin/data/ent/role"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuQuery is the builder for querying Menu entities.
type MenuQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	inters       []Interceptor
	predicates   []predicate.Menu
	withRoles    *RoleQuery
	withParent   *MenuQuery
	withChildren *MenuQuery
	withParams   *MenuParamQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MenuQuery builder.
func (mq *MenuQuery) Where(ps ...predicate.Menu) *MenuQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MenuQuery) Limit(limit int) *MenuQuery {
	mq.limit = &limit
	return mq
}

// Offset to start from.
func (mq *MenuQuery) Offset(offset int) *MenuQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MenuQuery) Unique(unique bool) *MenuQuery {
	mq.unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MenuQuery) Order(o ...OrderFunc) *MenuQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryRoles chains the current query on the "roles" edge.
func (mq *MenuQuery) QueryRoles() *RoleQuery {
	query := (&RoleClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, menu.RolesTable, menu.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (mq *MenuQuery) QueryParent() *MenuQuery {
	query := (&MenuClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, selector),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, menu.ParentTable, menu.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (mq *MenuQuery) QueryChildren() *MenuQuery {
	query := (&MenuClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, selector),
			sqlgraph.To(menu.Table, menu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, menu.ChildrenTable, menu.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParams chains the current query on the "params" edge.
func (mq *MenuQuery) QueryParams() *MenuParamQuery {
	query := (&MenuParamClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(menu.Table, menu.FieldID, selector),
			sqlgraph.To(menuparam.Table, menuparam.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, menu.ParamsTable, menu.ParamsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Menu entity from the query.
// Returns a *NotFoundError when no Menu was found.
func (mq *MenuQuery) First(ctx context.Context) (*Menu, error) {
	nodes, err := mq.Limit(1).All(newQueryContext(ctx, TypeMenu, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{menu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MenuQuery) FirstX(ctx context.Context) *Menu {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Menu ID from the query.
// Returns a *NotFoundError when no Menu ID was found.
func (mq *MenuQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = mq.Limit(1).IDs(newQueryContext(ctx, TypeMenu, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{menu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MenuQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Menu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Menu entity is found.
// Returns a *NotFoundError when no Menu entities are found.
func (mq *MenuQuery) Only(ctx context.Context) (*Menu, error) {
	nodes, err := mq.Limit(2).All(newQueryContext(ctx, TypeMenu, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{menu.Label}
	default:
		return nil, &NotSingularError{menu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MenuQuery) OnlyX(ctx context.Context) *Menu {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Menu ID in the query.
// Returns a *NotSingularError when more than one Menu ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MenuQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = mq.Limit(2).IDs(newQueryContext(ctx, TypeMenu, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{menu.Label}
	default:
		err = &NotSingularError{menu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MenuQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Menus.
func (mq *MenuQuery) All(ctx context.Context) ([]*Menu, error) {
	ctx = newQueryContext(ctx, TypeMenu, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Menu, *MenuQuery]()
	return withInterceptors[[]*Menu](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MenuQuery) AllX(ctx context.Context) []*Menu {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Menu IDs.
func (mq *MenuQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	ctx = newQueryContext(ctx, TypeMenu, "IDs")
	if err := mq.Select(menu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MenuQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MenuQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeMenu, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MenuQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MenuQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MenuQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeMenu, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MenuQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MenuQuery) Clone() *MenuQuery {
	if mq == nil {
		return nil
	}
	return &MenuQuery{
		config:       mq.config,
		limit:        mq.limit,
		offset:       mq.offset,
		order:        append([]OrderFunc{}, mq.order...),
		predicates:   append([]predicate.Menu{}, mq.predicates...),
		withRoles:    mq.withRoles.Clone(),
		withParent:   mq.withParent.Clone(),
		withChildren: mq.withChildren.Clone(),
		withParams:   mq.withParams.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MenuQuery) WithRoles(opts ...func(*RoleQuery)) *MenuQuery {
	query := (&RoleClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withRoles = query
	return mq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MenuQuery) WithParent(opts ...func(*MenuQuery)) *MenuQuery {
	query := (&MenuClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withParent = query
	return mq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MenuQuery) WithChildren(opts ...func(*MenuQuery)) *MenuQuery {
	query := (&MenuClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withChildren = query
	return mq
}

// WithParams tells the query-builder to eager-load the nodes that are connected to
// the "params" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MenuQuery) WithParams(opts ...func(*MenuParamQuery)) *MenuQuery {
	query := (&MenuParamClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withParams = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Menu.Query().
//		GroupBy(menu.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MenuQuery) GroupBy(field string, fields ...string) *MenuGroupBy {
	mq.fields = append([]string{field}, fields...)
	grbuild := &MenuGroupBy{build: mq}
	grbuild.flds = &mq.fields
	grbuild.label = menu.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Menu.Query().
//		Select(menu.FieldCreatedAt).
//		Scan(ctx, &v)
func (mq *MenuQuery) Select(fields ...string) *MenuSelect {
	mq.fields = append(mq.fields, fields...)
	sbuild := &MenuSelect{MenuQuery: mq}
	sbuild.label = menu.Label
	sbuild.flds, sbuild.scan = &mq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MenuSelect configured with the given aggregations.
func (mq *MenuQuery) Aggregate(fns ...AggregateFunc) *MenuSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MenuQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.fields {
		if !menu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MenuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Menu, error) {
	var (
		nodes       = []*Menu{}
		_spec       = mq.querySpec()
		loadedTypes = [4]bool{
			mq.withRoles != nil,
			mq.withParent != nil,
			mq.withChildren != nil,
			mq.withParams != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Menu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Menu{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withRoles; query != nil {
		if err := mq.loadRoles(ctx, query, nodes,
			func(n *Menu) { n.Edges.Roles = []*Role{} },
			func(n *Menu, e *Role) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withParent; query != nil {
		if err := mq.loadParent(ctx, query, nodes, nil,
			func(n *Menu, e *Menu) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := mq.withChildren; query != nil {
		if err := mq.loadChildren(ctx, query, nodes,
			func(n *Menu) { n.Edges.Children = []*Menu{} },
			func(n *Menu, e *Menu) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withParams; query != nil {
		if err := mq.loadParams(ctx, query, nodes,
			func(n *Menu) { n.Edges.Params = []*MenuParam{} },
			func(n *Menu, e *MenuParam) { n.Edges.Params = append(n.Edges.Params, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MenuQuery) loadRoles(ctx context.Context, query *RoleQuery, nodes []*Menu, init func(*Menu), assign func(*Menu, *Role)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*Menu)
	nids := make(map[uint64]map[*Menu]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(menu.RolesTable)
		s.Join(joinT).On(s.C(role.FieldID), joinT.C(menu.RolesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(menu.RolesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(menu.RolesPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := uint64(values[0].(*sql.NullInt64).Int64)
			inValue := uint64(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Menu]struct{}{byID[outValue]: {}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (mq *MenuQuery) loadParent(ctx context.Context, query *MenuQuery, nodes []*Menu, init func(*Menu), assign func(*Menu, *Menu)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*Menu)
	for i := range nodes {
		fk := nodes[i].ParentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(menu.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mq *MenuQuery) loadChildren(ctx context.Context, query *MenuQuery, nodes []*Menu, init func(*Menu), assign func(*Menu, *Menu)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Menu)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Menu(func(s *sql.Selector) {
		s.Where(sql.InValues(menu.ChildrenColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MenuQuery) loadParams(ctx context.Context, query *MenuParamQuery, nodes []*Menu, init func(*Menu), assign func(*Menu, *MenuParam)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint64]*Menu)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.MenuParam(func(s *sql.Selector) {
		s.Where(sql.InValues(menu.ParamsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.menu_params
		if fk == nil {
			return fmt.Errorf(`foreign-key "menu_params" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "menu_params" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: menu.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for i := range fields {
			if fields[i] != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(menu.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = menu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MenuGroupBy is the group-by builder for Menu entities.
type MenuGroupBy struct {
	selector
	build *MenuQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MenuGroupBy) Aggregate(fns ...AggregateFunc) *MenuGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MenuGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeMenu, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MenuQuery, *MenuGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MenuGroupBy) sqlScan(ctx context.Context, root *MenuQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MenuSelect is the builder for selecting fields of Menu entities.
type MenuSelect struct {
	*MenuQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MenuSelect) Aggregate(fns ...AggregateFunc) *MenuSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MenuSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeMenu, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MenuQuery, *MenuSelect](ctx, ms.MenuQuery, ms, ms.inters, v)
}

func (ms *MenuSelect) sqlScan(ctx context.Context, root *MenuQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
