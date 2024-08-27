package core

import (
	"context"
)

// AcceleratorKey is the key type of Instance in Context, exported for test.
type acceleratorKeyType int

const acceleratorKey acceleratorKeyType = 1

// FromContext returns an Instance from the given context, or nil if the context doesn't contain one.
func FromContext(ctx context.Context) *Instance {
	if s, ok := ctx.Value(acceleratorKey).(*Instance); ok {
		return s
	}
	return nil
}

// MustFromContext returns an Instance from the given context, or panics if not present.
func MustFromContext(ctx context.Context) *Instance {
	v := FromContext(ctx)
	if v == nil {
		panic("V is not in context.")
	}
	return v
}

/*
	toContext returns ctx from the given context, or creates an Instance if the context doesn't find that.

It is unsupported to use this function to create a context that is suitable to invoke Vdonkey's internal component
in third party code, you shouldn't use //go:linkname to alias of this function into your own package and
use this function in your third party code.

For third party code, usage enabled by creating a context to interact with Vdonkey's internal component is unsupported,
and may break at any time.
*/
func toContext(ctx context.Context, v *Instance) context.Context {
	if FromContext(ctx) != v {
		ctx = context.WithValue(ctx, acceleratorKey, v)
	}
	return ctx
}

/*
ToBackgroundDetachedContext create a detached context from another context
Internal API
*/
func ToBackgroundDetachedContext(ctx context.Context) context.Context {
	return &temporaryValueDelegationFix{context.Background(), ctx}
}

type temporaryValueDelegationFix struct {
	context.Context
	value context.Context
}

func (t *temporaryValueDelegationFix) Value(key interface{}) interface{} {
	return t.value.Value(key)
}
