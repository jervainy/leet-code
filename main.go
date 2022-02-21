package main

func strStr(haystack string, needle string) int {
	haylen, nelen := len(haystack), len(needle)
	if nelen == 0 {
		return 0
	}
	if nelen > haylen {
		return -1
	}
	for i := 0; i <= haylen - nelen; i++ {
		// 从右向左进行匹配
		var suffix []byte // 存储好后缀
		for j := nelen -1; j >= 0; j-- {
			if haystack[i + j] != needle[j] { // 出现了坏字符
				if len(suffix) == 0 { // 没有出现好后缀
					badChar := haystack[i + j] // 定义坏字符
					pos := j // 记录需要移动的位数
					for j0 := nelen -2; j0 >= 0; j0-- {
						if needle[j0] == badChar { // 模式串最右侧匹配到了坏字符
							pos = j - j0 -1 // for循环i会加1，此处需要减1
							break
						}
					}
					i += pos
					break
				} else { // 存在好后缀的情况下
					suffix = []byte{}
				}
			} else if j == 0 {
				return i
			} else {
				suffix = append(suffix, needle[j])
			}
		}
	}
	return -1
}

func main() {
	println(strStr("abbadcababacab", "babac"))
}
