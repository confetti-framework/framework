package middleware

import (
	"github.com/confetti-framework/framework/foundation/decorator/response_decorator"
	"github.com/confetti-framework/framework/inter"
)

type DecorateResponse struct{}

func (r DecorateResponse) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	decorators := request.App().Make("response_decorators").([]inter.ResponseDecorator)
	return response_decorator.Handler{Decorators: decorators}.Decorate(response)
}
