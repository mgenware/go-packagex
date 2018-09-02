package templatex

import (
	"io"
	"path/filepath"
	"text/template"
)

// View wraps a Go text/template.Template object, providing ease of use.
type View struct {
	// template holds the internal text/template Template object.
	template *template.Template
	file     string
	devMode  bool
}

// MustParseView loads a view from the given file, and panics if parsing failed.
func MustParseView(file string, devMode bool) *View {
	t := MustParseFile(file)
	return &View{template: t, file: file, devMode: devMode}
}

// MustParseViewFromDirectory joins the dir and the file arguments and calls MustParseView.
func MustParseViewFromDirectory(dir, file string, devMode bool) *View {
	return MustParseView(filepath.Join(dir, file), devMode)
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

// MustExecute calls Execute and panics if any error occurs.
func (view *View) MustExecute(wr io.Writer, data interface{}) {
	err := view.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

// MustExecuteToString calls ExecuteToString and panics if any error occurs.
func (view *View) MustExecuteToString(data interface{}) string {
	return MustExecuteToString(view.template, data)
}
