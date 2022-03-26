package section58

func lengthOfLastWord(s string) int {
	max, n := 0, len(s)
	for i := n -1; i >= 0; i-- {
		if max != 0 && s[i] == ' ' {
			return max
		}
		if s[i] != ' ' {
			max++
		}
	}
	return max
}
