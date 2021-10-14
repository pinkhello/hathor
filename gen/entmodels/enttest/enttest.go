// Code generated by entc, DO NOT EDIT.

package enttest

import (
	"context"

	"github.com/pinkhello/hathor/gen/entmodels"
	// required by schema hooks.
	_ "github.com/pinkhello/hathor/gen/entmodels/runtime"

	"entgo.io/ent/dialect/sql/schema"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...interface{})
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts        []entmodels.Option
		migrateOpts []schema.MigrateOption
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...entmodels.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// WithMigrateOptions forwards options to auto migration.
func WithMigrateOptions(opts ...schema.MigrateOption) Option {
	return func(o *options) {
		o.migrateOpts = append(o.migrateOpts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls entmodels.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *entmodels.Client {
	o := newOptions(opts)
	c, err := entmodels.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := c.Schema.Create(context.Background(), o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
	return c
}

// NewClient calls entmodels.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *entmodels.Client {
	o := newOptions(opts)
	c := entmodels.NewClient(o.opts...)
	if err := c.Schema.Create(context.Background(), o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
	return c
}