package main

import (
	"chat/ctx"
	"chat/http/server/router"
	"chat/pkg/http/engine"
)

func main() {
	serviceContext := ctx.NewDefaultServiceContext()
	router := route.NewRouter(serviceContext)
	ginEngine := engine.NewGinEngine(serviceContext.Conf().HttpPort(), router.Install)
	serviceContext.Start(ginEngine)
}
