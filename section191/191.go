package section191


func hammingWeight0(num uint32) int {
	res := 0
	for num > 0 {
		if num & 1 == 1 {
			res += 1
		}
		num >>= 1
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		num &= num - 1
		res += 1
	}
	return res
}

