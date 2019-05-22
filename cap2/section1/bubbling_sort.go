package section1

func BubblingSort(nums []int) []int {
	for k, v := range nums {
		if k == 0 {
			continue
		}
		cur := k
		for cur > 0 && nums[cur-1] > v {
			nums[cur], nums[cur-1] = nums[cur-1 ], nums[cur]
			cur--
		}
	}
	return nums
}

