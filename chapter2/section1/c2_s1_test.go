package section1

import (
	"testing"
)

func Test_BubblingSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 0}

	BubblingSort(nums)

	for k, v := range nums {
		if k != v {
			t.Error("failure")
			return
		}
	}
	t.Log("pass")
}

func Test_InsertionSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 0}

	InsertionSort(nums)

	for k, v := range nums {
		if k != v {
			t.Error("failure")
			return
		}
	}
	t.Log("pass")
}
