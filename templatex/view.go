package templatex

import (
	"bytes"
	"io"
	"text/template"
)

// GlobalDevMode determines if view is reloaded each time Execute is called, and defaults to false. Set it to true in development mode in order to applies template file changes without needing to restart Go server.
var GlobalDevMode = false

// View wraps a Go text/template.Template object, providing ease of use.
type View struct {
	// template holds the internal Template object from "text/template".
	template *template.Template
	file     string
	devMode  bool
}

// MustParseView loads a view from the given file, and panics if parsing failed.
func MustParseView(file string) *View {
	t := mustParseTemplate(file)
	return &View{template: t, file: file, devMode: GlobalDevMode}
}

// MustParseViewWithDevMode behaves like MustParseView, but allows user to override the devMode.
func MustParseViewWithDevMode(file string, devMode bool) *View {
	t := mustParseTemplate(file)
	return &View{template: t, file: file, devMode: devMode}
}

// Execute applies this view to the given data object.
func (view *View) Execute(wr io.Writer, data interface{}) error {
	if view.devMode {
		view.template = mustParseTemplate(view.file)
	}

	return view.template.Execute(wr, data)
}

// ExecuteToString applies this view to the given data object, and returns the result as a string.
func (view *View) ExecuteToString(data interface{}) (string, error) {
	buffer := &bytes.Buffer{}
	err := view.Execute(buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func mustParseTemplate(file string) *template.Template {
	return template.Must(template.New("T").ParseFiles(file))
}
