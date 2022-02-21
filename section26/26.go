package section26

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	k := 0
	for i := 0; i < numsLen; i++ {
		num := nums[i]
		if k == 0 || num != nums[k - 1] {
			nums[k] = num
			k += 1
		}
	}
	return k
}

