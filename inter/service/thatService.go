package inter

import (
	"github.com/wagertron/go-bones/model"
)

type ThatService interface {
	Get() (model.Thing, error)
	Put() (model.Thing, error)
	Delete() error
}
