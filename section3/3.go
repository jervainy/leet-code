package main

func lengthOfLongestSubstring0(s string) int {
	max := 0
	for i, n := 0, len(s); i < n; i++ {
		tmp, j := make(map[uint8]uint8), i
		for j < n {
			c := s[j]
			if _, ok := tmp[c]; ok {
				max = maxInt(max, j - i)
				break
			} else {
				tmp[c] = c
			}
			j++
		}
		if j == n {
			max = maxInt(max, j - i)
		}
	}
	return max
}




func lengthOfLongestSubstring(s string) int {
	max, t, n, left := 0, make(map[uint8]int), len(s), 0
	if n == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		c := s[i]
		if cIndex, ok := t[c]; ok {
			left = maxInt(left, cIndex + 1)
		}
		t[c] = i
		max = maxInt(max, i - left + 1)
	}
	return max
}



func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	println(lengthOfLongestSubstring("dvdf"))
}
