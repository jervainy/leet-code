package main

func mySqrt(x int) int {
	l, r, ans := 0, x, -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	println(mySqrt(9))
}
