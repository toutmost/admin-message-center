// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/toutmost/admin-message-center/ent/emailprovider"
)

// EmailProviderCreate is the builder for creating a EmailProvider entity.
type EmailProviderCreate struct {
	config
	mutation *EmailProviderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (epc *EmailProviderCreate) SetCreatedAt(t time.Time) *EmailProviderCreate {
	epc.mutation.SetCreatedAt(t)
	return epc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillableCreatedAt(t *time.Time) *EmailProviderCreate {
	if t != nil {
		epc.SetCreatedAt(*t)
	}
	return epc
}

// SetUpdatedAt sets the "updated_at" field.
func (epc *EmailProviderCreate) SetUpdatedAt(t time.Time) *EmailProviderCreate {
	epc.mutation.SetUpdatedAt(t)
	return epc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillableUpdatedAt(t *time.Time) *EmailProviderCreate {
	if t != nil {
		epc.SetUpdatedAt(*t)
	}
	return epc
}

// SetName sets the "name" field.
func (epc *EmailProviderCreate) SetName(s string) *EmailProviderCreate {
	epc.mutation.SetName(s)
	return epc
}

// SetAuthType sets the "auth_type" field.
func (epc *EmailProviderCreate) SetAuthType(u uint8) *EmailProviderCreate {
	epc.mutation.SetAuthType(u)
	return epc
}

// SetEmailAddr sets the "email_addr" field.
func (epc *EmailProviderCreate) SetEmailAddr(s string) *EmailProviderCreate {
	epc.mutation.SetEmailAddr(s)
	return epc
}

// SetPassword sets the "password" field.
func (epc *EmailProviderCreate) SetPassword(s string) *EmailProviderCreate {
	epc.mutation.SetPassword(s)
	return epc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillablePassword(s *string) *EmailProviderCreate {
	if s != nil {
		epc.SetPassword(*s)
	}
	return epc
}

// SetHostName sets the "host_name" field.
func (epc *EmailProviderCreate) SetHostName(s string) *EmailProviderCreate {
	epc.mutation.SetHostName(s)
	return epc
}

// SetIdentify sets the "identify" field.
func (epc *EmailProviderCreate) SetIdentify(s string) *EmailProviderCreate {
	epc.mutation.SetIdentify(s)
	return epc
}

// SetNillableIdentify sets the "identify" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillableIdentify(s *string) *EmailProviderCreate {
	if s != nil {
		epc.SetIdentify(*s)
	}
	return epc
}

// SetSecret sets the "secret" field.
func (epc *EmailProviderCreate) SetSecret(s string) *EmailProviderCreate {
	epc.mutation.SetSecret(s)
	return epc
}

// SetNillableSecret sets the "secret" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillableSecret(s *string) *EmailProviderCreate {
	if s != nil {
		epc.SetSecret(*s)
	}
	return epc
}

// SetPort sets the "port" field.
func (epc *EmailProviderCreate) SetPort(u uint32) *EmailProviderCreate {
	epc.mutation.SetPort(u)
	return epc
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillablePort(u *uint32) *EmailProviderCreate {
	if u != nil {
		epc.SetPort(*u)
	}
	return epc
}

// SetTLS sets the "tls" field.
func (epc *EmailProviderCreate) SetTLS(b bool) *EmailProviderCreate {
	epc.mutation.SetTLS(b)
	return epc
}

// SetNillableTLS sets the "tls" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillableTLS(b *bool) *EmailProviderCreate {
	if b != nil {
		epc.SetTLS(*b)
	}
	return epc
}

// SetIsDefault sets the "is_default" field.
func (epc *EmailProviderCreate) SetIsDefault(b bool) *EmailProviderCreate {
	epc.mutation.SetIsDefault(b)
	return epc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (epc *EmailProviderCreate) SetNillableIsDefault(b *bool) *EmailProviderCreate {
	if b != nil {
		epc.SetIsDefault(*b)
	}
	return epc
}

// SetID sets the "id" field.
func (epc *EmailProviderCreate) SetID(u uint64) *EmailProviderCreate {
	epc.mutation.SetID(u)
	return epc
}

// Mutation returns the EmailProviderMutation object of the builder.
func (epc *EmailProviderCreate) Mutation() *EmailProviderMutation {
	return epc.mutation
}

// Save creates the EmailProvider in the database.
func (epc *EmailProviderCreate) Save(ctx context.Context) (*EmailProvider, error) {
	epc.defaults()
	return withHooks(ctx, epc.sqlSave, epc.mutation, epc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (epc *EmailProviderCreate) SaveX(ctx context.Context) *EmailProvider {
	v, err := epc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epc *EmailProviderCreate) Exec(ctx context.Context) error {
	_, err := epc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epc *EmailProviderCreate) ExecX(ctx context.Context) {
	if err := epc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (epc *EmailProviderCreate) defaults() {
	if _, ok := epc.mutation.CreatedAt(); !ok {
		v := emailprovider.DefaultCreatedAt()
		epc.mutation.SetCreatedAt(v)
	}
	if _, ok := epc.mutation.UpdatedAt(); !ok {
		v := emailprovider.DefaultUpdatedAt()
		epc.mutation.SetUpdatedAt(v)
	}
	if _, ok := epc.mutation.TLS(); !ok {
		v := emailprovider.DefaultTLS
		epc.mutation.SetTLS(v)
	}
	if _, ok := epc.mutation.IsDefault(); !ok {
		v := emailprovider.DefaultIsDefault
		epc.mutation.SetIsDefault(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (epc *EmailProviderCreate) check() error {
	if _, ok := epc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EmailProvider.created_at"`)}
	}
	if _, ok := epc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EmailProvider.updated_at"`)}
	}
	if _, ok := epc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EmailProvider.name"`)}
	}
	if _, ok := epc.mutation.AuthType(); !ok {
		return &ValidationError{Name: "auth_type", err: errors.New(`ent: missing required field "EmailProvider.auth_type"`)}
	}
	if _, ok := epc.mutation.EmailAddr(); !ok {
		return &ValidationError{Name: "email_addr", err: errors.New(`ent: missing required field "EmailProvider.email_addr"`)}
	}
	if _, ok := epc.mutation.HostName(); !ok {
		return &ValidationError{Name: "host_name", err: errors.New(`ent: missing required field "EmailProvider.host_name"`)}
	}
	if _, ok := epc.mutation.TLS(); !ok {
		return &ValidationError{Name: "tls", err: errors.New(`ent: missing required field "EmailProvider.tls"`)}
	}
	if _, ok := epc.mutation.IsDefault(); !ok {
		return &ValidationError{Name: "is_default", err: errors.New(`ent: missing required field "EmailProvider.is_default"`)}
	}
	return nil
}

func (epc *EmailProviderCreate) sqlSave(ctx context.Context) (*EmailProvider, error) {
	if err := epc.check(); err != nil {
		return nil, err
	}
	_node, _spec := epc.createSpec()
	if err := sqlgraph.CreateNode(ctx, epc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	epc.mutation.id = &_node.ID
	epc.mutation.done = true
	return _node, nil
}

func (epc *EmailProviderCreate) createSpec() (*EmailProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailProvider{config: epc.config}
		_spec = sqlgraph.NewCreateSpec(emailprovider.Table, sqlgraph.NewFieldSpec(emailprovider.FieldID, field.TypeUint64))
	)
	if id, ok := epc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := epc.mutation.CreatedAt(); ok {
		_spec.SetField(emailprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := epc.mutation.UpdatedAt(); ok {
		_spec.SetField(emailprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := epc.mutation.Name(); ok {
		_spec.SetField(emailprovider.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := epc.mutation.AuthType(); ok {
		_spec.SetField(emailprovider.FieldAuthType, field.TypeUint8, value)
		_node.AuthType = value
	}
	if value, ok := epc.mutation.EmailAddr(); ok {
		_spec.SetField(emailprovider.FieldEmailAddr, field.TypeString, value)
		_node.EmailAddr = value
	}
	if value, ok := epc.mutation.Password(); ok {
		_spec.SetField(emailprovider.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := epc.mutation.HostName(); ok {
		_spec.SetField(emailprovider.FieldHostName, field.TypeString, value)
		_node.HostName = value
	}
	if value, ok := epc.mutation.Identify(); ok {
		_spec.SetField(emailprovider.FieldIdentify, field.TypeString, value)
		_node.Identify = value
	}
	if value, ok := epc.mutation.Secret(); ok {
		_spec.SetField(emailprovider.FieldSecret, field.TypeString, value)
		_node.Secret = value
	}
	if value, ok := epc.mutation.Port(); ok {
		_spec.SetField(emailprovider.FieldPort, field.TypeUint32, value)
		_node.Port = value
	}
	if value, ok := epc.mutation.TLS(); ok {
		_spec.SetField(emailprovider.FieldTLS, field.TypeBool, value)
		_node.TLS = value
	}
	if value, ok := epc.mutation.IsDefault(); ok {
		_spec.SetField(emailprovider.FieldIsDefault, field.TypeBool, value)
		_node.IsDefault = value
	}
	return _node, _spec
}

// EmailProviderCreateBulk is the builder for creating many EmailProvider entities in bulk.
type EmailProviderCreateBulk struct {
	config
	err      error
	builders []*EmailProviderCreate
}

// Save creates the EmailProvider entities in the database.
func (epcb *EmailProviderCreateBulk) Save(ctx context.Context) ([]*EmailProvider, error) {
	if epcb.err != nil {
		return nil, epcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(epcb.builders))
	nodes := make([]*EmailProvider, len(epcb.builders))
	mutators := make([]Mutator, len(epcb.builders))
	for i := range epcb.builders {
		func(i int, root context.Context) {
			builder := epcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailProviderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, epcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, epcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, epcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (epcb *EmailProviderCreateBulk) SaveX(ctx context.Context) []*EmailProvider {
	v, err := epcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (epcb *EmailProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := epcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (epcb *EmailProviderCreateBulk) ExecX(ctx context.Context) {
	if err := epcb.Exec(ctx); err != nil {
		panic(err)
	}
}
