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




func lengthOfLongestSubstring1(s string) int {
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


/**
滑动窗口方法
 */
func lengthOfLongestSubstring(s string) int {
	result, n, tmap := 0, len(s), make(map[uint8]int)
	if n == 0 {
		return 0
	}
	for left, right := 0, 0; left < n; {
		for ; right < n; right++ { // 移动右指针
			c := s[right]
			val := tmap[c]
			if val > 0 {
				result = maxInt(result, right - left)
				// 移动左指针
				break
			}
			tmap[c] = 1
		}
	}
	return result
}



func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	println(lengthOfLongestSubstring("pwwkew"))
}
