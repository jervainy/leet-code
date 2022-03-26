package main

func reverseWords(s string) string {
	var (
		sb = []byte{}
		start = -1
	)
	for i, n := 0, len(s); i < n; i++ {
		if s[i] == ' ' {
			if start != -1 {
				sb = reverse(sb, s, start, i - 1)
			}
			sb = append(sb, ' ')
			start = -1
		} else if start == -1 {
			start = i
			if i == n - 1 {
				sb = append(sb, s[i])
			}
		} else if i == n - 1 && start != -1 {
			sb = reverse(sb, s, start, i)
		}
	}
	return string(sb)
}

func reverse(sb []byte, s string, left, right int) []byte {
	for left <= right {
		sb = append(sb, s[right])
		right--
	}
	return sb
}

func main() {
	println(reverseWords("Let's take LeetCode contest"))
}
