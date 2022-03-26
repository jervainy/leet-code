package main

import "sort"

func merge0(nums1 []int, m int, nums2 []int, n int)  {
	for i := 0; i < n; i++ {
		nums1[m + i] = nums2[i]
	}
	sort.Ints(nums1)
}


func merge1(nums1 []int, m int, nums2 []int, n int)  {
	ans, l1, l2, i := make([]int, m + n), 0, 0, 0
	for l1 < m || l2 < n {
		if l1 >= m {
			ans[i] = nums2[l2]
			l2++
		} else if l2 >= n {
			ans[i] = nums1[l1]
			l1++
		} else {
			if nums1[l1] > nums2[l2] {
				ans[i] = nums2[l2]
				l2++
			} else {
				ans[i] = nums1[l1]
				l1++
			}
		}
		i++
	}
	copy(nums1, ans)
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
	l1, l2, i := m - 1, n - 1, m + n - 1
	for i >= 0 {
		if l1 < 0 {
			nums1[i] = nums2[l2]
			l2--
		} else if l2 < 0 {
			nums1[i] = nums1[l1]
			l1--
		} else {
			if nums1[l1] > nums2[l2] {
				nums1[i] = nums1[l1]
				l1--
			} else {
				nums1[i] = nums2[l2]
				l2--
			}
		}
		i--
	}
}





func main() {
	nums1, nums2 := []int{1,2,3,0,0,0}, []int{2,5,6}
	merge(nums1, 3, nums2, 3)
	println(nums1)
}
