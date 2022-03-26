package section66

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0 {
		return []int{0}
	}
	for i := n -1; i >= 0; i-- {
		d := digits[i]
		d += 1
		if d < 10 {
			digits[i] = d
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
