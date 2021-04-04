package templatex

import (
	"bytes"
	"os"
	"text/template"
)

// MustParse behaves like template.Parse, but panics if any error happens.
func MustParse(name, s string) *template.Template {
	return template.Must(template.New(name).Parse(s))
}

// MustParseFile reads a file from file parameter and calls MustParse. It panics when error occurs.
func MustParseFile(file string) *template.Template {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return MustParse("File "+file, string(content))
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

// MustExecuteToString calls ExecuteToString and panics if any error occurs.
func MustExecuteToString(t *template.Template, data interface{}) string {
	result, err := ExecuteToString(t, data)
	if err != nil {
		panic(err)
	}
	return result
}
