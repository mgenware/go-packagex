package templatex

import (
	"bytes"
	"text/template"
)

// MustParse behaves like template.Parse, but panics if any error happens.
func MustParse(s string) *template.Template {
	return template.Must(template.New("T").Parse(s))
}

// MustParseFromFile reads a file from file parameter and calls MustParse. It panics when error occurs.
func MustParseFromFile(file string) *template.Template {
	return template.Must(template.New("T").ParseFiles(file))
}

// ExecuteToString executes a template with specified data and returns a string.
func ExecuteToString(t *template.Template, data interface{}) (string, error) {
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
