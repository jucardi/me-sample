package helloworld

import (
	"github.com/jucardi/go-titan/net/errorx"
	"github.com/jucardi/go-titan/net/rest"
	"github.com/jucardi/go-titan/net/rest/router"
	. "github.com/jucardi/ms-sample/api/helloworld"
	"strings"
)

var (
	restRouterInstance = &restRouter{}
)

// AddRoutes registers routes for HelloworldApi to the API router.
func AddRoutes(router router.IRouter) {
	r1 := router.Version(ApiVersion).Group(ApiRoute)
	r1.GET("", restRouterInstance.getHelloWorldHandler)
	r1.PUT(RouteMessages, restRouterInstance.create)
	r1.GET(RouteMessagesParam, restRouterInstance.get)
	r1.PATCH(RouteMessages, restRouterInstance.update)
}

// Router definition for IApi
type restRouter struct {
}

func (r *restRouter) getHelloWorldHandler(c *rest.Context) {
	c.SendOrErr(Controller().WithCtx(c).GetHelloWorld())
}

func (r *restRouter) create(c *rest.Context) {
	req := &Message{}

	if err := c.Bind(req); err != nil {
		c.SendError(errorx.WrapBadRequest(err, "failed to deserialize request body"))
	} else {
		c.StatusOrErr(Controller().WithCtx(c).Create(req))
	}
}

func (r *restRouter) get(c *rest.Context) {
	var (
		fetcher  func(string) string
		paramVal string
	)

	// Param sessionID string
	if strings.Contains(RouteMessagesParam, ":name") {
		fetcher = c.Param
	} else {
		fetcher = c.Query
	}
	paramVal = fetcher("name")
	name := paramVal

	c.SendOrErr(Controller().WithCtx(c).Get(name))
}

func (r *restRouter) update(c *rest.Context) {
	req := &Message{}

	if err := c.Bind(req); err != nil {
		c.SendError(errorx.WrapBadRequest(err, "failed to deserialize request body"))
	} else {
		c.StatusOrErr(Controller().WithCtx(c).Update(req))
	}
}
