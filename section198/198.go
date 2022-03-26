package section198

func rob(nums []int) int {
	n, f1 := len(nums), nums[0]
	if n == 1 {
		return f1
	}
	f2 := max(f1, nums[1])
	for i := 2; i< n; i++ {
		tmp := max(f1 + nums[i], f2)
		f1 = f2
		f2 = tmp
	}
	return f2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

