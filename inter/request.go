package inter

import (
	"github.com/confetti-framework/framework/support"
	"net/http"
)

const RequestBodyDecoder = "request_body_decoder"

type Request interface {
	App() App
	SetApp(app App)
	Make(abstract interface{}) interface{}
	MakeE(abstract interface{}) (interface{}, error)
	Body() string
	SetBody(body string) Request
	Source() http.Request
	Header(key string) string
	Headers() http.Header
	Cookie(key string) string
	CookieE(key string) (string, error)
	File(key string) support.File
	FileE(key string) (support.File, error)
	Files(key string) []support.File
	FilesE(key string) ([]support.File, error)
	Method() string
	Path() string
	Url() string
	FullUrl() string
	Content(key ...string) support.Value
	ContentE(keyInput ...string) (support.Value, error)
	ContentOr(keys string, defaultValue interface{}) support.Value
	Parameter(key string) support.Value
	ParameterE(key string) (support.Value, error)
	ParameterOr(key string, defaultValue interface{}) support.Value
	SetUrlValues(vars map[string]string) Request
	Query(key string) support.Value
	QueryE(key string) (support.Value, error)
	QueryOr(key string, defaultValue interface{}) support.Value
	Route() Route
}
