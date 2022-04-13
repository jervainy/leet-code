package main

func majorityElement0(nums []int) int {
	n := len(nums)
	quickSort(nums, 0, n - 1)
	maxsum, sum, maxnum := 1, 1, nums[0]
	for i := 1; i < n; i++ {
		if nums[i] == nums[i - 1] {
			sum += 1
		} else {
			sum = 1
		}
		if sum > maxsum {
			maxsum = sum
			maxnum = nums[i]
		}
	}
	return maxnum
}

func majorityElement(nums []int) int {
	n := len(nums)
	quickSort(nums, 0, n - 1)
	return nums[n / 2]
}

func quickSort(nums []int, l, r int)  {
	key, i, j := nums[(l + r) / 2], l, r
	for i < j {
		for i < j && nums[i] < key {
			i += 1
		}
		for i < j && nums[j] > key {
			j -= 1
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		if nums[i] == key {
			j -= 1
		}
		if nums[j] == key {
			i += 1
		}
	}
	if i - l > 1 {
		quickSort(nums, l, i - 1)
	}
	if j - i > 1 {
		quickSort(nums, j + 1, r)
	}
}

func main() {
	nums := []int{2,2,1,3,1,1,4,1,1,5,1,1,6}
	majorityElement(nums)
	println(nums)
}
