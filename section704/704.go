package section704

func search(nums []int, target int) int {
	left, right, pos := 0, len(nums), -1
	for left <= right {
		mid := left + (right - left) >> 1
		if nums[mid] == target {
			pos = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid -1
		}
	}
	return pos
}

func main() {

}

