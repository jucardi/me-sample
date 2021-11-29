package api

import (
	"github.com/jucardi/go-titan/logx"
	"github.com/jucardi/go-titan/net/rest/router"
	"github.com/jucardi/ms-sample/server/api/helloworld"
)

func Start() {
	r := router.FromConfig()
	initRouter(r)

	logx.WithObj(
		r.Run(),
	).Fatal("Error occurred while starting the REST router")
}

func StartAdmin() {
	r := router.FromConfig()

	logx.WithObj(
		r.RunAdmin(),
	).Fatal("Error occurred while starting the REST router")
}

func initRouter(router router.IRouter) {
	// Add routes per domain under the currently defined ContextPath
	helloworld.AddRoutes(router)
}
