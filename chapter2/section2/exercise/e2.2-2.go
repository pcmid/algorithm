package exercise

func SelectionSort(nums []int) []int {

	getMinIndex := func(n []int) int {
		min := 0
		for k := range n {
			if n[min] > n[k] {
				min = k
			}
		}
		return min
	}

	for i := 0; i < len(nums)-1; i++ {
		minIndex := getMinIndex(nums[i:])
		if nums[i+minIndex] < nums[i] {
			nums[i], nums[i+minIndex] = nums[i+minIndex], nums[i]
		}
	}

	return nums
}
