package main


func rotate0(nums []int, k int)  {
	n := len(nums)
	r := n - k % n
	ans := append(nums[r:], nums[:r]...)
	for i, n := 0, len(nums); i < n; i++ {
		nums[i] = ans[i]
	}
}

func rotate1(nums []int, k int)  {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		index := (i + k) % n
		ans[index] = nums[i]
	}
	copy(nums, ans)
}




func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n - 1 -i] = a[n - 1 -i], a[i]
	}
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	println(nums)
}



func main() {
	rotate([]int{1,2,3,4,5,6,7}, 3)
}
