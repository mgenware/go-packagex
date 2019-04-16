package templatex

import (
	"testing"

	"github.com/mgenware/go-packagex/v5/test"
)

func TestExecuteToString(t *testing.T) {
	tpl := MustParse("v:{{.}}")
	got, err := ExecuteToString(tpl, "1")
	if err != nil {
		panic(err)
	}
	exp := "v:1"
	test.Compare(t, exp, got)
}

func TestMustParse(t *testing.T) {
	tpl := MustParse("{{.}}_{{html .}}")
	got, err := ExecuteToString(tpl, "<")
	if err != nil {
		panic(err)
	}
	exp := "<_&lt;"
	test.Compare(t, exp, got)
}
