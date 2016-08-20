package stringsx

import "testing"

func Test_SubString(t *testing.T) {
	res := "b"
	if SubString("abc", 1, 2) != res {
		t.Errorf("Expected %v", res)
	}

	res = " 一二"
	if SubString(" 一二三456", 0, 3) != res {
		t.Errorf("Expected %v", res)
	}
}

func Test_SubStringFromStart(t *testing.T) {
	res := "bcdef"
	if SubStringFromStart("abcdef", 1) != res {
		t.Errorf("Expected %v", res)
	}

	res = "三456"
	if SubStringFromStart(" 一二三456", 3) != res {
		t.Errorf("Expected %v", res)
	}
}

func Test_SubStringToEnd(t *testing.T) {
	res := "a"
	if SubStringToEnd("abcdef", 1) != res {
		t.Errorf("Expected %v", res)
	}

	res = " 一二"
	if SubStringToEnd(" 一二三456", 3) != res {
		t.Errorf("Expected %v", res)
	}
}

func Test_Truncate(t *testing.T) {
	res := "a"
	if Truncate("abcdef", 1) != res {
		t.Errorf("Expected %v", res)
	}

	res = " 一二"
	if Truncate(" 一二三456", 3) != res {
		t.Errorf("Expected %v", res)
	}

	res = "123一二三"
	if Truncate("123一二三", 20) != res {
		t.Errorf("Expected %v", res)
	}

	res = "123一二三"
	if Truncate("123一二三", 6) != res {
		t.Errorf("Expected %v", res)
	}
}

func Test_Reverse(t *testing.T) {
	res := "fedcba"
	if Reverse("abcdef") != res {
		t.Errorf("Expected %v", res)
	}

	res = "二一321"
	if Reverse("123一二") != res {
		t.Errorf("Expected %v", res)
	}

	res = "654三二一"
	if Reverse("一二三456") != res {
		t.Errorf("Expected %v", res)
	}
}

func Test_JoinAll(t *testing.T) {
	res := "1,2,abc,<nil>"
	if JoinAll([]interface{}{1, 2, "abc", nil}, ",") != res {
		t.Errorf("Expected %v", res)
	}

	res = ""
	if JoinAll([]interface{}{}, ",") != res {
		t.Errorf("Expected %v", res)
	}
}
