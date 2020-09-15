package templatex

import (
	"testing"

	"github.com/mgenware/go-packagex/v5/test"
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

func TestMustParse(t *testing.T) {
	tpl := MustParse("T", "{{.}}_{{html .}}")
	got, err := ExecuteToString(tpl, "<")
	if err != nil {
		panic(err)
	}
	exp := "<_&lt;"
	test.Assert(t, got, exp)
}
