package section67

func addBinary0(a string, b string) string {
	an, bn, i := len(a), len(b), 1
	var result []byte
	var carry byte = '0'
	for {
		ac, bc := charAt(a, an - i, an), charAt(b, bn - i, bn)
		if ac == '0' && bc == '0' && carry == '0' && i > an && i > bn {
			break
		}
		r, c := bitAdd(ac, bc, carry)
		carry = c
		result = append([]byte{r}, result...)
		i++
	}
	if len(result) == 0 {
		return "0"
	}
	return string(result)
}

func bitAdd(a, b, c byte) (byte, byte) {
	sum := (a - 48) + (b - 48) + (c - 48)
	if sum == 0 {
		return '0', '0'
	} else if sum == 1 {
		return '1', '0'
	} else if sum == 2 {
		return '0', '1'
	}
	return '1', '1'
}

func charAt(str string, index, n int) byte {
	if index >= n || index < 0 {
		return '0'
	}
	return str[index]
}

func main() {
	println(addBinary("100", "110010"))
	b := (0b10 & 0b10) << 1
	println(string(b))
}
