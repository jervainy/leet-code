package main


var (
	dx = []int{0,1,0,-1}
	dy = []int{1,0,-1,0}
)

func updateMatrix(mat [][]int) [][]int {
	var (
		m = len(mat)
		n = len(mat[0])
		t = make(map[int]int)
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			t[i * m + j] = bfs(i, j, m, n, mat)
		}
	}
	return mat
}

func bfs(i, j, m, n int, mat [][]int) int {
	var (
		queue []int
		list []int
		deep = 0
	)
	queue = append(queue, i * m + j)
	for len(queue) > 0 {
		x, y := queue[0] / m, queue[0] % m
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			kx, ky := dx[k] + x, dy[k] + y
			if kx >= 0 && kx < m && ky >= 0 && ky < n {
				if mat[kx][ky] == 0 {
					return deep
				} else {
					list = append(list, kx * m + ky)
				}
			}
		}
		if len(queue) == 0 {
			if len(list) == 0 {
				deep = 0
			} else {
				deep += 1
				queue = append(queue, list...)
				list = list[0:0]
			}
		}
	}

	return deep
}

func main() {
	mat := [][]int{
		{0,0,0},
		{0,1,0},
		{1,1,1},
	}
	updateMatrix(mat)
	println(mat)
}
