package main


func moveZeroes0(nums []int)  {
	for i, n := 0, len(nums); i < n; i++ {
		ZERO: if nums[i] == 0 {
			for j := i + 1; j < n; j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					i += 1
					goto ZERO
				}
			}
			return
		}
	}
}


func moveZeroes1(nums []int)  {
	for i, j, n := 0, 0, len(nums); i < n; i++ {
		if nums[i] != 0 {
			if nums[j] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j += 1
		}
	}
}

func moveZeroes(nums []int)  {
	j, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j += 1
		}
	}
	for j < n {
		nums[j] = 0
		j++
	}
}



func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	println(nums)
}
