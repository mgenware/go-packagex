package templatex

import (
	"io"
	"text/template"
)

var globalDevMode = false

// SetGlobalDevMode sets the global dev mode variable (defaults to false). If true, view will be always reloaded before being executed, so that template file changes apply without needing to restart Go server.
func SetGlobalDevMode(value bool) {
	globalDevMode = value
}

// View wraps a Go text/template.Template object, providing ease of use.
type View struct {
	// template holds the internal text/template Template object.
	template *template.Template
	file     string
	devMode  bool
}

// MustParseView loads a view from the given file, and panics if parsing failed.
func MustParseView(file string) *View {
	return MustParseViewWithDevMode(file, globalDevMode)
}

// MustParseViewWithDevMode behaves like MustParseView, but allows user to override the devMode.
func MustParseViewWithDevMode(file string, devMode bool) *View {
	t := MustParseFile(file)
	return &View{template: t, file: file, devMode: devMode}
}

// Execute applies this view to the given data object.
func (view *View) Execute(wr io.Writer, data interface{}) error {
	if view.devMode {
		view.template = MustParseFile(view.file)
	}
	return view.template.Execute(wr, data)
}

// ExecuteToString applies this view to the given data object, and returns the result as a string.
func (view *View) ExecuteToString(data interface{}) (string, error) {
	return ExecuteToString(view.template, data)
}
