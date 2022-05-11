package inter

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"io"
)

type ExitCode int

const (
	Success ExitCode = iota
	Failure
	Index
	Continue
)

const (
	Short = "short"
	Flag = "flag"
	Description = "description"
)

type ConsoleKernel interface {
	Handle() ExitCode
}

type Command interface {
	Name() string
	Description() string
	Handle(c Cli) ExitCode
}

type Cli interface {
	App() App
	Writer() io.Writer
	WriterErr() io.Writer
	Ask(label string) string
	Secret(label string) string
	Confirm(label string, defaultValue bool) bool
	Choice(label string, items ...string) string
	Line(format string, arguments ...interface{})
	Comment(format string, arguments ...interface{})
	Info(format string, arguments ...interface{})
	Error(format string, arguments ...interface{})
	Table() table.Writer
}
