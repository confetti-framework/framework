package response_decorator

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/framework/contract/inter"
)

type HttpStatus struct {
	ErrorDefault int
}

func (h HttpStatus) Decorate(response inter.Response) inter.Response {
	if err, ok := response.GetContent().(error); ok {
		status, ok := errors.FindStatus(err)
		if !ok && h.ErrorDefault != 0 {
			status = h.ErrorDefault
		}
		response.Status(status)
	}

	return response
}
