package stringsx

import (
	"testing"

	"github.com/mgenware/go-packagex/test"
)

func TestSubString(t *testing.T) {
	test.Compare(t, "b", SubString("abc", 1, 2))
	test.Compare(t, " 一二", SubString(" 一二三456", 0, 3))
}

func TestSubStringFromStart(t *testing.T) {
	test.Compare(t, "bcdef", SubStringFromStart("abcdef", 1))
	test.Compare(t, "三456", SubStringFromStart(" 一二三456", 3))
}

func TestSubStringToEnd(t *testing.T) {
	test.Compare(t, "a", SubStringToEnd("abcdef", 1))
	test.Compare(t, " 一二", SubStringToEnd(" 一二三456", 3))
}

func TestTruncate(t *testing.T) {
	test.Compare(t, "a", Truncate("abcdef", 1))
	test.Compare(t, " 一二", Truncate(" 一二三456", 3))
	test.Compare(t, "123一二三", Truncate("123一二三", 20))
	test.Compare(t, "123一二三", Truncate("123一二三", 6))
}

func TestReverse(t *testing.T) {
	test.Compare(t, "fedcba", Reverse("abcdef"))
	test.Compare(t, "二一321", Reverse("123一二"))
	test.Compare(t, "654三二一", Reverse("一二三456"))
}

func TestJoinAll(t *testing.T) {
	test.Compare(t, "1,2,abc,<nil>", JoinAll([]interface{}{1, 2, "abc", nil}, ","))
	test.Compare(t, "", JoinAll([]interface{}{}, ","))
}
