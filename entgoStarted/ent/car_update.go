// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"entgoStarted/ent/car"
	"entgoStarted/ent/predicate"
	"entgoStarted/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetModel sets the "model" field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.mutation.SetModel(s)
	return cu
}

// SetRegisteredAt sets the "registered_at" field.
func (cu *CarUpdate) SetRegisteredAt(t time.Time) *CarUpdate {
	cu.mutation.SetRegisteredAt(t)
	return cu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cu *CarUpdate) SetOwnerID(id int) *CarUpdate {
	cu.mutation.SetOwnerID(id)
	return cu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cu *CarUpdate) SetNillableOwnerID(id *int) *CarUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the "owner" edge to the User entity.
func (cu *CarUpdate) SetOwner(u *User) *CarUpdate {
	return cu.SetOwnerID(u.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cu *CarUpdate) Mutation() *CarMutation {
	return cu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cu *CarUpdate) ClearOwner() *CarUpdate {
	cu.mutation.ClearOwner()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: car.FieldModel,
		})
	}
	if value, ok := cu.mutation.RegisteredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: car.FieldRegisteredAt,
		})
	}
	if cu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetModel sets the "model" field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.mutation.SetModel(s)
	return cuo
}

// SetRegisteredAt sets the "registered_at" field.
func (cuo *CarUpdateOne) SetRegisteredAt(t time.Time) *CarUpdateOne {
	cuo.mutation.SetRegisteredAt(t)
	return cuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cuo *CarUpdateOne) SetOwnerID(id int) *CarUpdateOne {
	cuo.mutation.SetOwnerID(id)
	return cuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableOwnerID(id *int) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the "owner" edge to the User entity.
func (cuo *CarUpdateOne) SetOwner(u *User) *CarUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cuo *CarUpdateOne) Mutation() *CarMutation {
	return cuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (cuo *CarUpdateOne) ClearOwner() *CarUpdateOne {
	cuo.mutation.ClearOwner()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Car entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	var (
		err  error
		node *Car
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (_node *Car, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Model(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: car.FieldModel,
		})
	}
	if value, ok := cuo.mutation.RegisteredAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: car.FieldRegisteredAt,
		})
	}
	if cuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Car{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
