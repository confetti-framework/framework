package http_helper

import (
	"github.com/confetti-framework/framework/inter"
	"strings"
)

func IsMethod(request inter.Request, method string) bool {
	return request.Method() == strings.ToUpper(method)
}
