package section28

/**
BF算法 Brute Force
暴力
 */

func strStr_BF(haystack string, needle string) int {
	haylen, nelen := len(haystack), len(needle)
	if nelen == 0 {
		return 0
	}
	if nelen > haylen {
		return -1
	}
	for i := 0; i <= haylen - nelen; i++ {
		for j := 0; j < nelen; j++ {
			if haystack[i + j] != needle[j] {
				break
			} else if j == nelen -1 {
				return i
			}
		}
	}
	return -1
}
