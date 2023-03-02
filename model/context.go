package model

import (
	"context"
	"errors"
)

var ourContextKey = struct{}{} // always will be unique with empty struct, and not exported so GetOurValues will be necessary to call it

type OurContext struct {
	context.Context
}

type OurValues struct {
	SomeValue string
}

func NewOurContext(vals OurValues) context.Context {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, ourContextKey, vals)

	return ctx
}

func GetOurValues(ctx context.Context) (*OurValues, error) {
	val, ok := ctx.Value(ourContextKey).(OurValues)
	if !ok {
		return nil, GetOurValuesContextError{
			errors.New("error retrieving OurValues"),
		}
	}

	return &val, nil
}
