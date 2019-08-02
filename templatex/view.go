package templatex

import (
	"io"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

// View wraps a Go text/template.Template object, providing ease of use.
type View struct {
	// Use Template() to access the internal template data, which takes devMode into account.
	internalTemplate *template.Template
	file             string
	// Used in dev mode to detect file changes.
	lastModified time.Time
	devMode      bool
}

// MustParseView loads a view from the given file, and panics if parsing failed.
func MustParseView(file string, devMode bool) *View {
	t := MustParseFile(file)
	return &View{internalTemplate: t, file: file, devMode: devMode}
}

// MustParseViewFromDirectory joins the dir and the file arguments and calls MustParseView.
func MustParseViewFromDirectory(dir, file string, devMode bool) *View {
	return MustParseView(filepath.Join(dir, file), devMode)
}

// Execute applies this view to the given data object.
func (view *View) Execute(wr io.Writer, data interface{}) error {
	return view.Template().Execute(wr, data)
}

// ExecuteToString applies this view to the given data object, and returns the result as a string.
func (view *View) ExecuteToString(data interface{}) (string, error) {
	return ExecuteToString(view.Template(), data)
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
	return MustExecuteToString(view.Template(), data)
}

// Template returns the underlying template.Template.
func (view *View) Template() *template.Template {
	if view.devMode {
		// Get the last modified time
		stat, err := os.Stat(view.file)
		// Can panic in dev mode
		if err != nil {
			panic(err)
		}
		modTime := stat.ModTime()
		if !modTime.Equal(view.lastModified) {
			view.internalTemplate = MustParseFile(view.file)
			view.lastModified = modTime
		}
	}
	return view.internalTemplate
}
