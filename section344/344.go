package section344


func reverseString0(s []byte)  {
	for i, n := 0, len(s); i < n / 2; i++ {
		s[i], s[n - 1 - i] = s[n - 1 - i], s[i]
	}
}


func reverseString(s []byte)  {
	for left, right := 0, len(s) - 1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
