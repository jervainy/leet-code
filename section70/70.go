package main

func climbStairs0(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return climbStairs(n - 1) + climbStairs(n - 2)
}

func climbStairs(n int) int {
	a, b, sum := 0, 0, 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			a = 1
			sum = 1
		} else if i == 2 {
			b = 2
			sum = 2
		} else {
			sum = a + b
			a = b
			b = sum
		}
	}
	return sum
}

func main() {
	// 1134903170
	println(climbStairs(44))
}
