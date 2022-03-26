package section1137

func tribonacci(n int) int {
	n0, n1, n2, sum := 0, 1, 1, 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			sum = n0
		} else if i == 1 {
			sum = n1
		} else if i == 2 {
			sum = n2
		} else {
			sum = n0 + n1 + n2
			n0 = n1
			n1 = n2
			n2 = sum
		}
	}
	return sum
}
