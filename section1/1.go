package section1

func twoSum1(nums []int, target int) []int {
	var (
		result = []int{0, 0}
		n      = len(nums)
	)
	if n == 0 {
		return result
	}
	compare := false
	for i := 0; i < n; i++ {
		if compare == false {
			result[0] = i
			compare = true
			if i == n -1 {
				return result
			}
		} else {
			if nums[result[0]] + nums[i] == target {
				result[1] = i
				return result
			} else if i == n -1 {
				i = result[0]
				compare = false
			}
		}
	}
	return result
}


func twoSum(nums []int, target int) []int {
	m, n := make(map[int]int), len(nums)
	for i := 0; i < n; i++ {
		sub := target - nums[i]
		val, ok := m[sub]
		if ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}
