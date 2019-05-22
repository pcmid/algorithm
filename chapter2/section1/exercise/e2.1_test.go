package exercise

import (
	"testing"
)

func Test_Seek(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6}

	if Seek(nums, 6) != 6 {
		t.Error("seeked failure")
	} else {
		t.Log("seeked pass")
	}

	if Seek(nums, 7) != 0 {
		t.Error("unseeked failure")
	} else {
		t.Log("unseeked pass")
	}
}

func Test_BinaryAdd(t *testing.T) {
	b1 := [8]int{1, 0, 0, 1, 1, 1, 0, 1}
	b2 := [8]int{1, 0, 0, 1, 1, 1, 0, 1}

	res := BinaryAdd(b1[:], b2[:])

	ans := []int{1, 0, 0, 1, 1, 1, 0, 1, 0}
	for k := range res {
		if res[k] != ans[k] {
			t.Error("failure")
			return
		}
	}
	t.Log("pass")
}
