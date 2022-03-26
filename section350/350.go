package section350

func intersect(nums1 []int, nums2 []int) []int {
	m, ans := map[int]int{}, []int{}
	for i, n := 0, len(nums1); i < n; i++ {
		if v, ok := m[nums1[i]]; ok {
			m[nums1[i]] = v + 1
		} else {
			m[nums1[i]] = 1
		}
	}

	for i, n := 0, len(nums2); i < n; i++ {
		if v, ok := m[nums2[i]]; ok && v > 0 {
			ans = append(ans, nums2[i])
			m[nums2[i]] = v - 1
		}
	}
	return ans
}
