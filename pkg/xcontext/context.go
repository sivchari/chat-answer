package xcontext

import (
	"context"
)

type UserID struct{}

type keyConstraint interface {
	UserID
}

type valueConstraint interface {
	~string
}

type key[T keyConstraint] struct{}

func WithValue[k keyConstraint, v valueConstraint](ctx context.Context, val v) context.Context {
	return context.WithValue(ctx, key[k]{}, val)
}

func Value[k keyConstraint, v valueConstraint](ctx context.Context) (v, bool) {
	val, ok := ctx.Value(key[k]{}).(v)
	return val, ok
}
