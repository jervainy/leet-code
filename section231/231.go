package main

func isPowerOfTwo0(n int) bool {
	return n > 0 && n & (n - 1) == 0
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n & (-n) == n
}

func main() {
	println(5 >> 1)
	println(isPowerOfTwo(16))
}
