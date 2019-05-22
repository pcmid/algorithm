package exercise

func Seek(nums []int, v int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] == v {
			return i
		}
	}
	return 0
}
