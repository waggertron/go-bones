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

func (s *MySvc) Get() (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (s *MySvc) Put() (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (s *MySvc) Delete() error {
	return nil
}
