package templatex

import (
	"html/template"
	"testing"

	"github.com/mgenware/go-packagex/v6/test"
)

func TestExecuteToString(t *testing.T) {
	tpl := MustParse("T", "v:{{.}}")
	got, err := ExecuteToString(tpl, "1")
	if err != nil {
		panic(err)
	}
	exp := "v:1"
	test.Assert(t, got, exp)
}

type TestTemplateData struct {
	A string
	B template.HTML
}

func TestMustParse(t *testing.T) {
	d := TestTemplateData{A: "<", B: template.HTML("<")}
	tpl := MustParse("T", "{{.A}}_{{.B}}")
	got, err := ExecuteToString(tpl, d)
	if err != nil {
		panic(err)
	}
	exp := "&lt;_<"
	test.Assert(t, got, exp)
}
