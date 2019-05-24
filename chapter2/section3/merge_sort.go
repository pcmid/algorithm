package section3

func MergeSort(nums []int) {

	numsL := nums[:len(nums)/2]
	numsR := nums[len(nums)/2:]

	if len(numsL) > 1 {
		MergeSort(numsL)
	}
	if len(numsR) > 1 {
		MergeSort(numsR)
	}
	Merge(nums)

}

func Merge(nums []int) {
	var numLS []int
	var numRS []int
	for i := range nums[:len(nums)/2] {
		numLS = append(numLS, nums[i])

	}
	for i := range nums[len(nums)/2:] {
		numRS = append(numRS, nums[len(nums)/2+i])

	}

	{
		l := 0
		r := 0

		for i := 0; i < len(nums); i++ {
			if l >= len(numLS) {
				nums[i] = numRS[r]
				r++
				continue
			}
			if r >= len(numRS) {
				nums[i] = numLS[l]
				l++
				continue
			}
			if numLS[l] <= numRS[r] {
				nums[i] = numLS[r]
				l++
			} else {
				nums[i] = numRS[r]
				r++
			}
		}
	}
}
