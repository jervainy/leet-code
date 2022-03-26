package main

func main() {
	a := []int{}
	a = append(a, 1)
	a = append(a, 2)
	a = a[0:1]
	a = append(a, 3)
	print(a)
}
