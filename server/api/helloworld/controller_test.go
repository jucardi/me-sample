package helloworld

import (
	"net/http"
	"testing"

	. "github.com/jucardi/go-testx/testx"
	"github.com/jucardi/go-titan/configx"
	"github.com/jucardi/go-titan/net/rest/router"
	"github.com/jucardi/go-titan/utils/testutils"
	"github.com/jucardi/ms-sample/api/helloworld"
	"github.com/jucardi/ms-sample/server/repository"
)

func init() {
	// Loads the test configuration
	_ = configx.FromFile(testutils.DefaultTestConfigPath())

	// Initializes the database connection
	repository.InitRepositories()

	// Initializes the test router
	testutils.Prepare(router.Bare(), AddRoutes)
}

func TestHelloRoute(t *testing.T) {
	var testRoute = testutils.ContextURI(helloworld.ApiVersion, helloworld.ApiRoute)
	Convey("When calling "+testRoute, t, func() {
		req, _ := http.NewRequest(http.MethodGet, testRoute, nil)
		resp := testutils.Serve(req)

		Convey("Expecting a status of 200 and hello message", t, func() {

			message := &helloworld.Message{}

			ShouldEqual(resp.Code, 200)
			ShouldNotError(resp.Unmarshal(message))
			ShouldEqual("Hello World! Welcome to a Go Microservice.", message.Message)
		})
	})

}
