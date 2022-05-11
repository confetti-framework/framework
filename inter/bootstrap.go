package inter

type Bootstrap interface {
	Bootstrap(app Container) Container
}

type RouteDecorator interface {
	Decorate(route Route) Route
}

type ResponseDecorator interface {
	Decorate(response Response) Response
}
