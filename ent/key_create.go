// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/riadafridishibly/vault-app/ent/key"
)

// KeyCreate is the builder for creating a Key entity.
type KeyCreate struct {
	config
	mutation *KeyMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (kc *KeyCreate) SetType(s string) *KeyCreate {
	kc.mutation.SetType(s)
	return kc
}

// SetPublicKey sets the "public_key" field.
func (kc *KeyCreate) SetPublicKey(s string) *KeyCreate {
	kc.mutation.SetPublicKey(s)
	return kc
}

// SetPrivateKey sets the "private_key" field.
func (kc *KeyCreate) SetPrivateKey(s string) *KeyCreate {
	kc.mutation.SetPrivateKey(s)
	return kc
}

// SetCreatedAt sets the "created_at" field.
func (kc *KeyCreate) SetCreatedAt(t time.Time) *KeyCreate {
	kc.mutation.SetCreatedAt(t)
	return kc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kc *KeyCreate) SetNillableCreatedAt(t *time.Time) *KeyCreate {
	if t != nil {
		kc.SetCreatedAt(*t)
	}
	return kc
}

// SetUpdatedAt sets the "updated_at" field.
func (kc *KeyCreate) SetUpdatedAt(t time.Time) *KeyCreate {
	kc.mutation.SetUpdatedAt(t)
	return kc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kc *KeyCreate) SetNillableUpdatedAt(t *time.Time) *KeyCreate {
	if t != nil {
		kc.SetUpdatedAt(*t)
	}
	return kc
}

// Mutation returns the KeyMutation object of the builder.
func (kc *KeyCreate) Mutation() *KeyMutation {
	return kc.mutation
}

// Save creates the Key in the database.
func (kc *KeyCreate) Save(ctx context.Context) (*Key, error) {
	var (
		err  error
		node *Key
	)
	kc.defaults()
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			if node, err = kc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			if kc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, kc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Key)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from KeyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KeyCreate) SaveX(ctx context.Context) *Key {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KeyCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KeyCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KeyCreate) defaults() {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		v := key.DefaultCreatedAt()
		kc.mutation.SetCreatedAt(v)
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		v := key.DefaultUpdatedAt()
		kc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KeyCreate) check() error {
	if _, ok := kc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Key.type"`)}
	}
	if _, ok := kc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "public_key", err: errors.New(`ent: missing required field "Key.public_key"`)}
	}
	if _, ok := kc.mutation.PrivateKey(); !ok {
		return &ValidationError{Name: "private_key", err: errors.New(`ent: missing required field "Key.private_key"`)}
	}
	if _, ok := kc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Key.created_at"`)}
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Key.updated_at"`)}
	}
	return nil
}

func (kc *KeyCreate) sqlSave(ctx context.Context) (*Key, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (kc *KeyCreate) createSpec() (*Key, *sqlgraph.CreateSpec) {
	var (
		_node = &Key{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: key.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: key.FieldID,
			},
		}
	)
	if value, ok := kc.mutation.GetType(); ok {
		_spec.SetField(key.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := kc.mutation.PublicKey(); ok {
		_spec.SetField(key.FieldPublicKey, field.TypeString, value)
		_node.PublicKey = value
	}
	if value, ok := kc.mutation.PrivateKey(); ok {
		_spec.SetField(key.FieldPrivateKey, field.TypeString, value)
		_node.PrivateKey = value
	}
	if value, ok := kc.mutation.CreatedAt(); ok {
		_spec.SetField(key.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := kc.mutation.UpdatedAt(); ok {
		_spec.SetField(key.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// KeyCreateBulk is the builder for creating many Key entities in bulk.
type KeyCreateBulk struct {
	config
	builders []*KeyCreate
}

// Save creates the Key entities in the database.
func (kcb *KeyCreateBulk) Save(ctx context.Context) ([]*Key, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Key, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KeyCreateBulk) SaveX(ctx context.Context) []*Key {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KeyCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KeyCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}