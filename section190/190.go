package main

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}
	var (
		res uint32 = 0
		count uint32 = 32
	)
	for num > 0 {
		b := num & 1
		res = res << 1 + b
		num >>= 1
		count--
	}
	return res << count
}

func main() {
	reverseBits(43261596)
}
