package inter

type MapMethodRoutes map[string][]Route

type RouteCollection interface {
	Push(route Route) RouteCollection
	Merge(routeCollections RouteCollection) RouteCollection
	All() []Route
	Match(request Request) Route
	Where(parameter, regex string) RouteCollection
	WhereMulti(constraints map[string]string) RouteCollection
	Domain(domain string) RouteCollection
	Prefix(prefix string) RouteCollection
	Name(name string) RouteCollection
	Middleware(...HttpMiddleware) RouteCollection
	WithoutMiddleware(...HttpMiddleware) RouteCollection
}
