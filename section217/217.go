package section217


func containsDuplicate(nums []int) bool {
	c := make(map[int]int)
	for i, n := 0, len(nums); i < n; i++ {
		if _, ok := c[nums[i]]; ok {
			return true
		}
		c[nums[i]] = nums[i]
	}
	return false
}
