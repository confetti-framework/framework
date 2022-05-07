package routing

import (
	"github.com/confetti-framework/framework/contract/inter"
	"github.com/confetti-framework/framework/foundation"
	"github.com/confetti-framework/framework/foundation/http"
	"github.com/confetti-framework/framework/foundation/test/mock"
)

func emptyController() func(request inter.Request) inter.Response {
	return func(request inter.Request) inter.Response { return nil }
}

func newRequest(options http.Options) inter.Request {
	app := foundation.NewApp()
	app.Bind("outcome_html_encoders", mock.HtmlEncoders)
	app.Bind("response_decorators", []inter.ResponseDecorator{})
	options.App = app
	return http.NewRequest(options)
}
