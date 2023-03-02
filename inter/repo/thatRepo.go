package inter

import (
	"context"

	"github.com/wagertron/go-bones/model"
)

type ThatRepo interface {
	Get(id int) (*model.Thing, error)
	Put(item model.Thing) (*model.Thing, error)
	Delete(id int) error
}
type ThatCompoundRepo interface {
	Get(ctx context.Context, id int) (model.Thing, error)
	Put(ctx context.Context, item model.Thing) (model.Thing, error)
	Delete(ctx context.Context, id int) error
}
