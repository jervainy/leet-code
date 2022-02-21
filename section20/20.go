package section20

func isValid(s string) bool {
	var (
		n = len(s)
		stack []byte
		pos = 0
		m = map[byte]byte{')': '(', ']': '[', '}': '{'}
	)
	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			pos++
		}
		if c == ')' || c == ']' || c == '}' {
			if pos == 0 || stack == nil {
				return false
			}
			if val := m[c]; val == stack[pos - 1] {
				pos--
				stack = stack[:pos]
			} else {
				return false
			}
		}
	}
	return pos == 0
}


