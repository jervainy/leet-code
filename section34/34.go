package main

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 || nums[0] > target || nums[n - 1] < target {
		return []int{-1, -1}
	}
	if n == 1 && nums[0] == target {
		return []int{0, 0}
	}
	left, right, retIndex := 0, n - 1, -1
	for left <= right {
		index := (left + right) / 2
		if nums[index] == target {
			retIndex = index
			break
		} else if nums[index] > target {
			right = index - 1
		} else if nums[index] < target {
			left = index + 1
		}
	}
	if retIndex == -1 {
		return []int{-1, -1}
	}
	var res = []int{retIndex, retIndex}
	for i := 1; retIndex - i >= 0 && nums[retIndex - i] == target; i++ {
		res[0] = retIndex - i
	}
	for i := 1; retIndex + i < n && nums[retIndex + i] == target; i++ {
		res[1] = retIndex + i
	}
	return res
}

func main() {
	searchRange([]int{1, 4}, 4)
}
