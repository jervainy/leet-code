package section53

import "math"

func maxSubArray0(nums []int) int {
	n, sum, max := len(nums), 0, 0
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	for i :=0; i < n; i++ {
		num := nums[i]
		if i == 0 {
			sum = num
			max = sum
		} else if sum < 0 && num > sum {
			sum = num
			max = int(math.Max(float64(sum), float64(max)))
		} else {
			if num < 0 {
				max = int(math.Max(float64(sum), float64(max)))
			}
			sum += num
		}
	}
	return int(math.Max(float64(sum), float64(max)))
}


func maxSubArray1(nums []int) int {
	max, n := nums[0], len(nums)
	for i := 1; i < n; i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}


// ----------------- 分治的方式解决 -------------------
func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if (l == r) {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}



