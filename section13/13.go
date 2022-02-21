package section13

func romanToInt(s string) int {
	result := 0
	if s == "" {
		return result
	}
	prev := -1
	for i := len(s) -1; i >= 0; i-- {
		c := s[i]
		num := 0
		if c == 'M' {
			num = 1000
		} else if c == 'D' {
			num = 500
		} else if c == 'C' {
			num = 100
		} else if c == 'L' {
			num = 50
		} else if c == 'X' {
			num = 10
		} else if c == 'V' {
			num = 5
		} else if c == 'I' {
			num = 1
		}
		if num >= prev {
			result += num
		} else {
			result -= num
		}
		prev = num
	}
	return result
}
