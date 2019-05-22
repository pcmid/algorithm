package exercise

func BinaryAdd(b1 []int, b2 []int) []int {
	res := make([]int, len(b1)+1)

	for i := len(b1) - 1; i >= 0; i-- {
		n := b1[i] + b2[i] + res[i+1]
		if n > 1 {
			res[i] = 1
		}
		res[i+1] = n % 2
	}
	return res
}
