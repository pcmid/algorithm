package exercise

import "testing"

func Test_SelectionSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1, 0}

	for k, v := range SelectionSort(nums) {
		if k != v {
			t.Error("failure")
			return
		}
	}
	t.Log("pass")
}
