package templatex

import "testing"

func TestExecuteToString(t *testing.T) {
	tpl := MustParse("v:{{.}}")
	got := ExecuteToString(tpl, "1")
	exp := "v:1"
	if got != exp {
		t.Errorf("Expected %v, got %v", exp, got)
	}
}

func TestMustParse(t *testing.T) {
	tpl := MustParse("{{.}}_{{html .}}")
	got := ExecuteToString(tpl, "<")
	exp := "<_&lt;"
	if got != exp {
		t.Errorf("Expected %v, got %v", exp, got)
	}
}
