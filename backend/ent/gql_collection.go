// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *InstanceQuery) CollectFields(ctx context.Context, satisfies ...string) *InstanceQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		i = i.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return i
}

func (i *InstanceQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *InstanceQuery {
	return i
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SlideQuery) CollectFields(ctx context.Context, satisfies ...string) *SlideQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *SlideQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *SlideQuery {
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
