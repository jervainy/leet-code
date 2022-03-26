package main

func rotate0(nums []int, k int)  {
	n := len(nums)
	ans := make([]int, n)
	coe := k / n + 1
	for i := 0; i < n; i++ {
		index := (n - k + i + coe * n) % n
		ans[i] = nums[index]
	}
	for i := 0; i < n; i++ {
		nums[i] = ans[i]
	}
}

func rotate(nums []int, k int)  {
	n := len(nums)
	r := n - k % n
	ans := append(nums[r:], nums[:r]...)
	for i, n := 0, len(nums); i < n; i++ {
		nums[i] = ans[i]
	}
}


func main() {
	nums := []int{1,2,3,4,5,6,7}
	a := nums[:]
	println(a)
}


