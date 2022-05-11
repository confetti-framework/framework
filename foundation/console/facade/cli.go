package facade

import (
	"fmt"
	"github.com/confetti-framework/framework/inter"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
	"io"
	"os"
	"time"
)

type cli struct {
	app       inter.App
	writer    io.Writer
	writerErr io.Writer
	reader    io.ReadCloser
}

var tableStyle = table.Style{
	Name: "Confetti",
	Box: table.BoxStyle{
		BottomLeft:       " ",
		BottomRight:      " ",
		BottomSeparator:  " ",
		EmptySeparator:   " ",
		Left:             " ",
		LeftSeparator:    " ",
		MiddleHorizontal: " ",
		MiddleSeparator:  " ",
		MiddleVertical:   " ",
		PaddingLeft:      " ",
		PaddingRight:     " ",
		PageSeparator:    "\n",
		Right:            " ",
		RightSeparator:   " ",
		TopLeft:          " ",
		TopRight:         " ",
		TopSeparator:     " ",
		UnfinishedRow:    " ~",
	},
	Color:  table.ColorOptionsDefault,
	Format: table.FormatOptionsDefault,
	HTML:   table.DefaultHTMLOptions,
	Options: table.Options{
		DrawBorder:      true,
		SeparateColumns: true,
		SeparateFooter:  true,
		SeparateHeader:  true,
		SeparateRows:    false,
	},
	Title: table.TitleOptionsBright,
}

func NewCli(app inter.App, writers ...io.Writer) *cli {
	c := &cli{app: app}
	setWriters(writers, c)

	return c
}

func NewCliByReadersAndWriter(
	app inter.App,
	reader io.ReadCloser,
	writer io.Writer,
	writerErr io.Writer,
) *cli {
	c := NewCli(app, writer, writerErr)
	c.reader = reader
	if c.reader == nil {
		c.reader = os.Stdin
	}

	return c
}

func (c *cli) App() inter.App {
	return c.app
}

func (c *cli) Writer() io.Writer {
	return c.writer
}

func (c *cli) WriterErr() io.Writer {
	return c.writerErr
}

func (c cli) Ask(label string) string {
	prompt := promptui.Prompt{
		Label:  label,
		Stdin:  c.reader,
		Stdout: writeCloser{writer: c.writer},
	}
	result, err := prompt.Run()
	if err != nil {
		_, _ = fmt.Fprintf(c.writer, "Prompt failed %v\n", err)
		os.Exit(int(inter.Failure))
	}
	return result
}

func (c cli) Secret(label string) string {
	prompt := promptui.Prompt{
		Label:  label,
		Mask:   '*',
		Stdin:  c.reader,
		Stdout: writeCloser{writer: c.writer},
	}
	result, err := prompt.Run()
	if err != nil {
		_, _ = fmt.Fprintf(c.writer, "Prompt failed %v\n", err)
		os.Exit(int(inter.Failure))
	}
	return result
}

func (c cli) Confirm(label string, defaultValue bool) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
		Stdin:     c.reader,
		Stdout:    writeCloser{writer: c.writer},
	}
	if defaultValue {
		prompt.Default = "y"
	}
	result, err := prompt.Run()
	if err != nil {
		return defaultValue
	}

	switch result {
	case "y", "Y", "yes", "Yes", "YES":
		return true
	case "":
		return defaultValue
	default:
		return false
	}
}

func (c cli) Choice(label string, items ...string) string {
	prompt := promptui.Select{
		Label:  label,
		Items:  items,
		Stdin:  c.reader,
		Stdout: writeCloser{writer: c.writer},
	}
	_, selected, err := prompt.Run()
	if err != nil {
		_, _ = fmt.Fprintf(c.writer, "Prompt failed %v\n", err)
		os.Exit(int(inter.Failure))
	}
	return selected
}

func (c *cli) Info(format string, arguments ...interface{}) {
	c.printColor("\033[32m", format, arguments)
}

func (c *cli) Error(format string, arguments ...interface{}) {
	_, err := fmt.Fprintf(c.writerErr, "\033[31m"+format+"\033[39m\n", arguments...)
	if err != nil {
		panic(err)
	}
}

func (c *cli) Line(format string, arguments ...interface{}) {
	c.printColor("\033[39m", format, arguments)
}

func (c *cli) Comment(format string, arguments ...interface{}) {
	c.printColor("\u001B[30;1m", format, arguments)
}

func (c cli) Table() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(c.Writer())
	t.SetStyle(tableStyle)
	return t
}

func (c *cli) ProgressBar(max int64, description ...string) *progressbar.ProgressBar {
	desc := ""
	if len(description) > 0 {
		desc = description[0]
	}
	bar := progressbar.NewOptions64(
		max,
		progressbar.OptionSetDescription(desc),
		progressbar.OptionSetWriter(c.writer),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(c.writer, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)
	bar.RenderBlank()
	return bar
}

type writeCloser struct {
	writer io.Writer
}

func (w writeCloser) Write(b []byte) (n int, err error) {
	return w.writer.Write(b)
}

func (w writeCloser) Close() error {
	return nil
}

func (c cli) printColor(color string, f string, a []interface{}) {
	_, err := fmt.Fprintf(c.writer, color+f+"\033[39m\n", a...)
	if err != nil {
		panic(err)
	}
}

func setWriters(writers []io.Writer, c *cli) {
	switch len(writers) {
	case 1:
		c.writer = writers[0]
		// Use the normal writer for err
		c.writerErr = writers[0]
	case 2:
		c.writer = writers[0]
		c.writerErr = writers[1]
		if c.writerErr == nil {
			// Use the normal writer for err
			c.writerErr = writers[0]
		}
	}

	// If nothing is given or the writer is nil, use the writer of the os.
	if c.writer == nil {
		c.writer = os.Stdout
	}
	if c.writerErr == nil {
		c.writerErr = os.Stderr
	}
}
