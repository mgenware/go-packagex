package templatex

import "testing"

func TestExecuteToString(t *testing.T) {
	tpl := MustParse("v:{{.}}")
	got, err := ExecuteToString(tpl, "1")
	if err != nil {
		panic(err)
	}
	exp := "v:1"
	if got != exp {
		t.Errorf("Expected %v, got %v", exp, got)
	}
}

func TestMustParse(t *testing.T) {
	tpl := MustParse("{{.}}_{{html .}}")
	got, err := ExecuteToString(tpl, "<")
	if err != nil {
		panic(err)
	}
	exp := "<_&lt;"
	if got != exp {
		t.Errorf("Expected %v, got %v", exp, got)
	}
}
