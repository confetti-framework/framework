package inter

type Next = PipeHolder

type HttpMiddleware interface {
	Handle(request Request, next Next) Response
}
