package route_decorator

import (
	"github.com/confetti-framework/framework/inter"
	"strings"
)

type UriPrefixSlash struct{}

func (o UriPrefixSlash) Decorate(route inter.Route) inter.Route {
	uri := route.Uri()

	if !strings.HasPrefix(uri, "/") {
		route.SetUri("/" + uri)
	}

	return route
}
