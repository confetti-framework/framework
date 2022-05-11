package inter

import "html/template"

type View interface {
	// Template returns the content of a template
	Template() string
}

type TemplateBuilder = func(template *template.Template) (*template.Template, error)
