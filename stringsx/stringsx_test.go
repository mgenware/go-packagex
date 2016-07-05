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
}
