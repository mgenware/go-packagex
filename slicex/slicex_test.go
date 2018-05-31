package slicex

import "testing"

func TestDeepEqualInt(t *testing.T) {
	if DeepEqualInt([]int{1, 2}, []int{1, 2}) == false {
		t.Fatal("Failed at [1,2], [1,2]")
	}
	if DeepEqualInt([]int{}, []int{}) == false {
		t.Fatal("Failed at [], []")
	}
	if DeepEqualInt(nil, nil) == false {
		t.Fatal("Failed at nil, nil")
	}
	if DeepEqualInt([]int{}, nil) == true {
		t.Fatal("Failed at [], nil")
	}
	if DeepEqualInt([]int{1, 2}, []int{}) == true {
		t.Fatal("Failed at [1,2], []")
	}
}

func TestDeepEqualUint(t *testing.T) {
	if DeepEqualUint([]uint{1, 2}, []uint{1, 2}) == false {
		t.Fatal("Failed at [1,2], [1,2]")
	}
	if DeepEqualUint([]uint{}, []uint{}) == false {
		t.Fatal("Failed at [], []")
	}
	if DeepEqualUint(nil, nil) == false {
		t.Fatal("Failed at nil, nil")
	}
	if DeepEqualUint([]uint{}, nil) == true {
		t.Fatal("Failed at [], nil")
	}
	if DeepEqualUint([]uint{1, 2}, []uint{}) == true {
		t.Fatal("Failed at [1,2], []")
	}
}

func TestDeepEqualInt64(t *testing.T) {
	if DeepEqualInt64([]int64{1, 2}, []int64{1, 2}) == false {
		t.Fatal("Failed at [1,2], [1,2]")
	}
	if DeepEqualInt64([]int64{}, []int64{}) == false {
		t.Fatal("Failed at [], []")
	}
	if DeepEqualInt64(nil, nil) == false {
		t.Fatal("Failed at nil, nil")
	}
	if DeepEqualInt64([]int64{}, nil) == true {
		t.Fatal("Failed at [], nil")
	}
	if DeepEqualInt64([]int64{1, 2}, []int64{}) == true {
		t.Fatal("Failed at [1,2], []")
	}
}

func TestDeepEqualUint64(t *testing.T) {
	if DeepEqualUint64([]uint64{1, 2}, []uint64{1, 2}) == false {
		t.Fatal("Failed at [1,2], [1,2]")
	}
	if DeepEqualUint64([]uint64{}, []uint64{}) == false {
		t.Fatal("Failed at [], []")
	}
	if DeepEqualUint64(nil, nil) == false {
		t.Fatal("Failed at nil, nil")
	}
	if DeepEqualUint64([]uint64{}, nil) == true {
		t.Fatal("Failed at [], nil")
	}
	if DeepEqualUint64([]uint64{1, 2}, []uint64{}) == true {
		t.Fatal("Failed at [1,2], []")
	}
}

func TestDeepEqualByte(t *testing.T) {
	if DeepEqualByte([]byte{1, 2}, []byte{1, 2}) == false {
		t.Fatal("Failed at [1,2], [1,2]")
	}
	if DeepEqualByte([]byte{}, []byte{}) == false {
		t.Fatal("Failed at [], []")
	}
	if DeepEqualByte(nil, nil) == false {
		t.Fatal("Failed at nil, nil")
	}
	if DeepEqualByte([]byte{}, nil) == true {
		t.Fatal("Failed at [], nil")
	}
	if DeepEqualByte([]byte{1, 2}, []byte{}) == true {
		t.Fatal("Failed at [1,2], []")
	}
}
