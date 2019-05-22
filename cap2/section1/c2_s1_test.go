package section1

import (
	"testing"
)

func Test_BubblingSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 0}

	for k, v := range BubblingSort(nums) {
		if k != v {
			t.Error("failure")
			return
		}
	}
	t.Log("pass")
}

func Test_InsertionSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 0}

	for k, v := range InsertionSort(nums) {
		if k != v {
			t.Error("failure")
			return
		}
	}
	t.Log("pass")
}
