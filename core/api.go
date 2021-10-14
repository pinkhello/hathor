package core

import (
	"fmt"

	"github.com/pinkhello/hathor/core/conf"
	"github.com/pinkhello/hathor/core/internal/handler"
	"github.com/pinkhello/hathor/core/internal/svc"
	"github.com/tal-tech/go-zero/rest"
)

func BuildServer(env string) (*rest.Server, *svc.ServiceContext) {
	conf := conf.LoadConfig(env)
	ctx := svc.NewServiceContext(conf)
	server := rest.MustNewServer(conf.Api)
	defer server.Stop()
	handler.RegisterHandlers(server, ctx)
	return server, ctx
}

func Start(env string) {
	server, ctx := BuildServer(env)
	fmt.Printf("Start Http Server at %s:%d.", ctx.Config.Api.Host, ctx.Config.Api.Port)
	server.Start()
}
