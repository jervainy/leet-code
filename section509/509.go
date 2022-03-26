package section509


func fib(n int) int {
	n0, n1, sum := 0, 1, 0
	for i := 0; i <= n; i++ {
		if i == 0 {
			sum = n0
		} else if i == 1 {
			sum = n1
		} else {
			sum = n0 + n1
			n0 = n1
			n1 = sum
		}
	}
	return sum
}
