package inter

type Route interface {
	Uri() string
	SetUri(url string) Route
	Method() string
	Controller() Controller
	SetPrefix(prefix string) Route
	SetDestination(destination string) Route
	SetStatus(status int) Route
	RouteOptions() RouteOptions
	Constraint() map[string]string
	SetConstraint(parameter string, regex string) Route
	Domain() string
	SetDomain(domain string) Route
	Name() string
	SetName(name string) Route
	Named(pattern ...string) bool
	Middleware() []HttpMiddleware
	SetMiddleware(middlewares []HttpMiddleware) Route
	SetExcludeMiddleware(middlewares []HttpMiddleware) Route
}

type RouteOptions interface {
	Prefixes() []string
	Status() int
}
