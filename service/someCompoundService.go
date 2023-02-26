package service

import (
	inter "github.com/wagertron/go-bones/inter/repo"
	"github.com/wagertron/go-bones/model"
)

type MyCompoundSvcCfg struct {
	Name     string
	ThisRepo inter.ThatRepo
}

type MyCompoundSvc struct {
	name string
	r    inter.ThatRepo
}

func NewMyCompoundSvc(c MyCompoundSvcCfg) *MyCompoundSvc {
	s := MyCompoundSvc{
		name: c.Name,
		r:    c.ThisRepo,
	}

	return &s
}

func (s *MyCompoundSvc) Get() (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (s *MyCompoundSvc) Put() (model.Thing, error) {
	t := model.Thing{}
	return t, nil
}

func (s *MyCompoundSvc) Delete() error {
	return nil
}
