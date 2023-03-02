package service

import (
	"context"

	inter "github.com/wagertron/go-bones/inter/repo"
	"github.com/wagertron/go-bones/model"
)

type MySvcCfg struct {
	Name     string
	ThisRepo inter.ThatRepo
	Ctx      context.Context
}

type MySvc struct {
	name string
	r    inter.ThatRepo
}

func NewMySvc(c MySvcCfg) *MySvc {
	s := MySvc{
		name: c.Name,
		r:    c.ThisRepo,
	}

	return &s
}

func (s *MySvc) Get(ctx context.Context, id int) (model.Thing, error) {
	item, err := s.r.Get(id)
	if err != nil {
		return model.Thing{}, model.NewGetError(err)
	}

	return *item, nil
}

func (s *MySvc) Put(ctx context.Context, updatedItem model.Thing) (model.Thing, error) {
	finalItem, err := s.r.Put(updatedItem)
	if err != nil {
		return model.Thing{}, model.NewPutError(err)
	}

	return *finalItem, nil
}

func (s *MySvc) Delete(ctx context.Context, id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return model.NewDeleteError(err)
	}

	return nil
}
