package routing

import (
	"github.com/confetti-framework/framework/foundation/http/outcome"
	"github.com/confetti-framework/framework/inter"
)

func redirectController(request inter.Request) inter.Response {
	rawRoute := request.App().Make("route")
	if rawRoute == nil {
		panic("no route found in request")
	}
	route := rawRoute.(Route)
	options := route.routeOptions
	return outcome.Redirect(options.destination, options.status)
}
