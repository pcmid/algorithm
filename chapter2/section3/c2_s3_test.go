package section3

import "testing"

func Test_Merge(t *testing.T) {
	nums := []int{1, 3, 5, 2, 4, 6, 7}

	Merge(nums)

	for i := range nums {
		if nums[i] != i+1 {
			t.Error("failure")
		}
	}
}

func Test_MergeSort(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	MergeSort(nums)

	for i := range nums {
		if nums[i] != i+1 {
			t.Error("failure")
		}
	}
}
