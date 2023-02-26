package repo

import (
	"context"

	"github.com/wagertron/go-bones/model"
)

type MyRepoCfg struct {
	Name string
	Ctx  context.Context
}

type MyRepo struct {
	name string
}

func NewMyRepo(c MyRepoCfg) *MyRepo {
	r := MyRepo{
		name: c.Name,
	}
	return &r
}

func (r *MyRepo) Get() (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (r *MyRepo) Put() (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (r *MyRepo) Delete() error {
	return nil
}
