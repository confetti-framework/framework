package route_decorator

import "github.com/confetti-framework/framework/contract/inter"

func Decorate(route inter.Route, decorators []inter.RouteDecorator) inter.Route {
	for _, decorator := range decorators {
		route = decorator.Decorate(route)
	}

	return route
}
