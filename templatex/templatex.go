package templatex

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

// MustParse behaves like template.Parse, but panics if any error happens.
func MustParse(s string) *template.Template {
	return template.Must(template.New("T").Parse(s))
}

// MustParseFromFile reads a file from file parameter and calls MustParse. It panics when error occurs.
func MustParseFromFile(file string) *template.Template {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return MustParse(string(bytes))
}

// ExecuteToString executes a template with specified data and returns a string.
func ExecuteToString(t *template.Template, data interface{}) string {
	buffer := &bytes.Buffer{}
	err := t.Execute(buffer, data)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
