package section994


func orangesRotting(grid [][]int) int {
	m, n, to, fo := len(grid), len(grid[0]), 0, 0
	for i := 0; i < m * n; i++ {
		r, c := i / n, i % n
		if grid[r][c] == 2 { // 腐烂的橘子
			fo = 2
		}
	}
	return to
}


