package stringsx

import (
	"testing"

	"github.com/mgenware/go-packagex/v6/test"
)

func TestSubString(t *testing.T) {
	test.Assert(t, SubString("abc", 1, 2), "b")
	test.Assert(t, SubString(" 一二三456", 0, 3), " 一二")
}

func TestSubStringFromStart(t *testing.T) {
	test.Assert(t, SubStringFromStart("abcdef", 1), "bcdef")
	test.Assert(t, SubStringFromStart(" 一二三456", 3), "三456")
}

func TestSubStringToEnd(t *testing.T) {
	test.Assert(t, SubStringToEnd("abcdef", 1), "a")
	test.Assert(t, SubStringToEnd(" 一二三456", 3), " 一二")
}

func TestTruncate(t *testing.T) {
	test.Assert(t, Truncate("abcdef", 1), "a")
	test.Assert(t, Truncate(" 一二三456", 3), " 一二")
	test.Assert(t, Truncate("123一二三", 20), "123一二三")
	test.Assert(t, Truncate("123一二三", 6), "123一二三")
}

func TestReverse(t *testing.T) {
	test.Assert(t, Reverse("abcdef"), "fedcba")
	test.Assert(t, Reverse("123一二"), "二一321")
	test.Assert(t, Reverse("一二三456"), "654三二一")
}

func TestJoinAll(t *testing.T) {
	test.Assert(t, JoinAll([]interface{}{1, 2, "abc", nil}, ","), "1,2,abc,<nil>")
	test.Assert(t, JoinAll([]interface{}{}, ","), "")
}
