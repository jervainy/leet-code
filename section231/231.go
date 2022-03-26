package main

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	for n >= 2 {
		if n == 2 {
			return true
		}
		n = n >> 1
	}
	return false
}

func main() {
	println(5 >> 1)
	println(isPowerOfTwo(16))
}
