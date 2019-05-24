package section1

func InsertionSort(nums []int) {
	for k, v := range nums {
		if k == 0 {
			continue
		}
		last := k - 1
		for last >= 0 && nums[last] > v {
			nums[last+1] = nums[last]
			last--
		}
		nums[last+1] = v
	}
}
