package model

import (
	"context"
	"errors"
)

const ourContextKey = "OUR_CONTEXT_VALUE"

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
