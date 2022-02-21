package section35

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right, pos := 0, n - 1, n
	for left <= right {
		if nums[left] >= target {
			return left
		} else if nums[right] == target {
			return right
		} else if nums[right] < target {
			return right + 1
		}
		mid := (right - left) >> 1 + left
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			pos = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return pos
}

func main() {
	searchInsert([]int{1,3,5,6}, 7)
}

