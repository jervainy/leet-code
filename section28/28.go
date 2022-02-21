package section28

func strStr0(haystack string, needle string) int {
	strLen, needleLen := len(haystack), len(needle)
	if needleLen == 0 {
		return 0
	}
	if needleLen > strLen {
		return -1
	}
	fc := needle[0]
	pos, nextPos := -1, -1
	for i := 0; i < strLen; i++ {
		c := haystack[i]
		if pos != -1 || c == fc {
			if pos == -1 {
				pos = i
				if needleLen == 1 {
					return pos
				}
				if i == strLen -1 {
					return -1
				}
				continue
			} else if c == fc && nextPos == -1 {
				nextPos = i
			}

			nIndex := i - pos
			nc := needle[nIndex]
			if c == nc {
				if nIndex == needleLen -1 {
					return pos
				}
				if i == strLen -1 {
					return -1
				}
			} else {
				pos = -1
				if nextPos != -1 {
					i = nextPos - 1
					nextPos = -1
				}
			}
		}
	}
	return pos
}

