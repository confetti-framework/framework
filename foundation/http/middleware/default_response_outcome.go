package middleware

import (
	"github.com/confetti-framework/framework/inter"
)

type DefaultResponseOutcome struct {
	Outcome func(content interface{}) inter.Response
}

func (r DefaultResponseOutcome) Handle(request inter.Request, next inter.Next) inter.Response {
	request.App().Bind("default_response_outcome", r.Outcome)
	return next(request)
}
