package section167

func twoSum0(numbers []int, target int) []int {
	for i ,n := 0, len(numbers); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if numbers[i] + numbers[j] == target {
				return []int{ i + 1, j + 1}
			}
		}
	}
	return []int{}
}


func twoSum1(numbers []int, target int) []int {
	tmp := make(map[int]int)
	for i ,n := 0, len(numbers); i < n; i++ {
		num := numbers[i]
		if tmpIndex, ok := tmp[target - num]; ok {
			return []int{tmpIndex + 1, i + 1}
		}
		tmp[num] = i
	}
	return []int{}
}


func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}


