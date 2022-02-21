package section14

import "bytes"

func longestCommonPrefix(strs []string) string {
	arrLen := len(strs)
	if arrLen == 0 {
		return ""
	} else if arrLen == 1 {
		return strs[0]
	}
	sb := bytes.Buffer{}
	index := 0
	for {
		if index >= len(strs[0]) {
			return sb.String()
		}
		c := strs[0][index]
		outter := true
		for i := 1; i < arrLen; i++ {
			if index >= len(strs[i])  {
				return sb.String()
			}
			if c != strs[i][index] {
				return sb.String()
			}
			outter = false
		}
		if outter {
			break
		}
		sb.WriteByte(c)
		index += 1
	}
	return sb.String()
}
