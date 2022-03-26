package main

func convertToTitle(columnNumber int) string {
	result, n := "", 26
	for columnNumber > 0 {
		columnNumber--
		a := columnNumber % n
		result = string('A' + byte(a)) + result
		columnNumber /= n
	}
	return result
}

func main() {
	println(convertToTitle(25))
}