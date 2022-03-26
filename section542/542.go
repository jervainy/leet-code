package main

import "math"

var (
	dx = []int{0,1,0,-1}
	dy = []int{1,0,-1,0}
)

type Point struct {
	x, y int
}

func updateMatrix0(mat [][]int) [][]int {
	var (
		m = len(mat)
		n = len(mat[0])
		dist = make([][]int, m)
		seen = make([][]bool, m)
		st []Point
	)

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		seen[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dist[i][j] = 0
				seen[i][j] = true
				st = append(st, Point{x: i, y: j})
			}
		}
	}

	for len(st) > 0 {
		p := st[0]
		st = st[1:]
		for i := 0; i < 4; i++ {
			px, py := p.x + dx[i], p.y + dy[i]
			if px >= 0 && px < m && py >= 0 && py < n && !seen[px][py] {
				seen[px][py] = true
				dist[px][py] = dist[p.x][p.y] + 1
				st = append(st, Point{x: px, y: py})
			}
		}
	}

	return dist
}










func updateMatrix(mat [][]int) [][]int {
	var (
		m = len(mat)
		n = len(mat[0])
		dist = make([][]int, m)
	)

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt / 2
			}
		}
	}

	// 右上
	for i := 0; i < m; i++ {
		for j := n - 1; j  >= 0; j-- {
			if i - 1 >= 0 {
				dist[i][j] = minInt(dist[i][j], dist[i - 1][j] + 1)
			}
			if j + 1 < n  {
				dist[i][j] = minInt(dist[i][j], dist[i][j + 1] + 1)
			}
		}
	}

	// 右下
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i - 1 >= 0 {
				dist[i][j] = minInt(dist[i][j], dist[i - 1][j] + 1)
			}
			if j - 1 >= n  {
				dist[i][j] = minInt(dist[i][j], dist[i - 1][j] + 1)
			}
		}
	}

	// 左上
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i + 1 < m {
				dist[i][j] = minInt(dist[i][j], dist[i + 1][j] + 1)
			}
			if j + 1 < n  {
				dist[i][j] = minInt(dist[i][j], dist[i][j + 1] + 1)
			}
		}
	}

	// 左下
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i + 1 < m {
				dist[i][j] = minInt(dist[i][j], dist[i + 1][j] + 1)
			}
			if j - 1 >= 0  {
				dist[i][j] = minInt(dist[i][j], dist[i][j - 1] + 1)
			}
		}
	}

	return dist
}

func minInt(x, y int) int {
	if x >= y {
		return y
	}
	return x
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
