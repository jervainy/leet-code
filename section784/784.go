package main


func letterCasePermutation(s string) []string {
	var (
		res []string
		dfs func(str string, path []byte, start int)
	)
	dfs = func(str string, path []byte, start int) {
		var tmp []byte
		c := str[start]
		tmp = append(tmp, c)
		t := caseTrans(c)
		if c != t {
			tmp = append(tmp, t)
		}
		for i, n := 0, len(tmp); i < n; i++ {
			path = append(path, tmp[i])
			if start < len(str) - 1 {
				dfs(str, path, start + 1)
			} else {
				res = append(res, string(path))
			}
			path = path[:len(path) - 1]
		}
	}
	dfs(s, []byte{}, 0)
	return res
}

func caseTrans(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + byte(32)
	}
	if c >= 'a' && c <= 'z' {
		return c - byte(32)
	}
	return c
}

func main() {
	letterCasePermutation("a1b2")
}
