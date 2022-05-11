package inter

type HttpKernel interface {
	Handle(request Request) Response
	RecoverFromMiddlewarePanic(recover interface{}) Response
}
