package main




func permute(nums []int) [][]int {
	var (
		res [][]int
		dfs func(arr, path []int)
	)
	dfs = func(arr, path []int) {
		for i, n := 0, len(arr); i < n; i++ {
			path = append(path, arr[i])
			newArr := sliceOne(arr, i)
			if len(newArr) == 0 {
				tmp := make([]int, len(path))
				copy(tmp, path)
				res = append(res, tmp)
			} else {
				dfs(newArr, path)
			}
			path = path[:len(path) - 1]
		}
	}

	dfs(nums, []int{})
	return res
}

func sliceOne(a []int, b int) []int {
	if len(a) == 0 {
		return a
	}
	if b == 0 {
		return a[1:]
	}
	if b >= len(a) - 1 {
		return a[:len(a) - 1]
	}
	var res []int
	res = append(res, a[:b]...)
	res = append(res, a[b+1:]...)
	return res
}

func main() {
	permute([]int{1,2,3})
}

