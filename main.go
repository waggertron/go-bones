package main

import (
	interRepo "github.com/wagertron/go-bones/inter/repo"
	interSvc "github.com/wagertron/go-bones/inter/service"
	interTrans "github.com/wagertron/go-bones/inter/transport"
	"github.com/wagertron/go-bones/repo"
	"github.com/wagertron/go-bones/service"
	"github.com/wagertron/go-bones/transport"
)

var (
	rCfg repo.MyRepoCfg
	r    interRepo.ThatRepo

	sCfg service.MySvcCfg
	s    interSvc.ThatService

	csCfg service.MyCompoundSvcCfg
	cs    interSvc.ThatService

	tCfg transport.HttpServerCfg
	t    interTrans.Server
)

func init() {
	// simple service
	rCfg = repo.MyRepoCfg{}
	r := repo.NewMyRepo(rCfg)
	or := repo.NewMyRepo(rCfg)

	sCfg = service.MySvcCfg{
		Name:          "serviceName",
		ThisRepo:      r,
		ThisOtherRepo: or,
	}
	s = service.NewMySvc(sCfg)

	// compound service (compound service repo = simple service)
	csCfg = service.MyCompoundSvcCfg{
		Name:     "compoundServiceName",
		ThisRepo: s,
	}
	cs = service.NewMyCompoundSvc(csCfg)

	tCfg = transport.HttpServerCfg{
		Name: "serverName",
		Port: ":9000",
	}
	t = transport.NewHttpServer(tCfg)

}

func main() {
	t.Start()
}
