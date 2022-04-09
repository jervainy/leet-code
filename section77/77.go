package main

import (
	"fmt"
	"unsafe"
)

func combine0(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if len(temp) + (n - cur + 1) < k {
			return
		}
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp) - 1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}





func combine(n int, k int) [][]int {
	var ans [][]int
	var dfs func(int, *[]int)
	dfs = func(cur int, path *[]int) {
		if len(*path) < k {
			for i := cur; i >= k; i-- {
				if len(*path) < k {
					*path = append(*path, i)
					dfs(cur - 1, path)
				}
			}
		}
	}

	for i := n; i >= k; i-- {
		var path []int
		path = append(path, i)
		dfs(i - 1, &path)
		ans = append(ans, path)
	}

	return ans
}


func main() {
	s := []byte{1,2,3,4}
	fmt.Printf("arrayA : %p , %v \n", &s[0], s[0])
	ptr := (*byte)(unsafe.Pointer(&s[0]))
	fmt.Printf("arrayA : %p , %v \n", ptr, *ptr)
}


