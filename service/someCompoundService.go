package service

import (
	"context"

	interRepo "github.com/wagertron/go-bones/inter/repo"
	interSvc "github.com/wagertron/go-bones/inter/service"
	"github.com/wagertron/go-bones/model"
)

type MyCompoundSvcCfg struct {
	Name     string
	ThisRepo interSvc.ThatService
}

type MyCompoundSvc struct {
	name string
	r    interRepo.ThatCompoundRepo
}

func NewMyCompoundSvc(c MyCompoundSvcCfg) *MyCompoundSvc {
	s := MyCompoundSvc{
		name: c.Name,
		r:    c.ThisRepo,
	}

	return &s
}

func (s *MyCompoundSvc) Get(ctx context.Context, ID int) (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (s *MyCompoundSvc) Put(ctx context.Context, updatedItem model.Thing) (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (s *MyCompoundSvc) Delete(ctx context.Context, ID int) error {
	return nil
}
