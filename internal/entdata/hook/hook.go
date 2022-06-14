// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"yola/internal/entdata"
)

// The MovieSourceFunc type is an adapter to allow the use of ordinary
// function as MovieSource mutator.
type MovieSourceFunc func(context.Context, *entdata.MovieSourceMutation) (entdata.Value, error)

// Mutate calls f(ctx, m).
func (f MovieSourceFunc) Mutate(ctx context.Context, m entdata.Mutation) (entdata.Value, error) {
	mv, ok := m.(*entdata.MovieSourceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *entdata.MovieSourceMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, entdata.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entdata.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entdata.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m entdata.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op entdata.Op) Condition {
	return func(_ context.Context, m entdata.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entdata.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entdata.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entdata.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk entdata.Hook, cond Condition) entdata.Hook {
	return func(next entdata.Mutator) entdata.Mutator {
		return entdata.MutateFunc(func(ctx context.Context, m entdata.Mutation) (entdata.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, entdata.Delete|entdata.Create)
//
func On(hk entdata.Hook, op entdata.Op) entdata.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, entdata.Update|entdata.UpdateOne)
//
func Unless(hk entdata.Hook, op entdata.Op) entdata.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) entdata.Hook {
	return func(entdata.Mutator) entdata.Mutator {
		return entdata.MutateFunc(func(context.Context, entdata.Mutation) (entdata.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entdata.Hook {
//		return []entdata.Hook{
//			Reject(entdata.Delete|entdata.Update),
//		}
//	}
//
func Reject(op entdata.Op) entdata.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entdata.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entdata.Hook) Chain {
	return Chain{append([]entdata.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entdata.Hook {
	return func(mutator entdata.Mutator) entdata.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entdata.Hook) Chain {
	newHooks := make([]entdata.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}