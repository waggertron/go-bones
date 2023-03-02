package inter

import (
	"context"

	"github.com/wagertron/go-bones/model"
)

type Server interface {
	Start() (context.CancelFunc, error)
	Stop() error
}

type Handler interface {
	Handle(event model.Event) (model.EventResponse, error)
}
