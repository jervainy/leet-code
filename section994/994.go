package section994

var (
	dx = []int{0,1,0,-1}
	dy = []int{1,0,-1,0}
)

type Point struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	var (
		m = len(grid)
		n = len(grid[0])
		seen = make([][]bool, m)
		st []Point
		t = -1
		s = 0
	)

	for i := 0; i < m; i++ {
		seen[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 { // 空格
				grid[i][j] = -1
			} else if grid[i][j] == 2 { // 腐烂的橘子
				grid[i][j] = 0
				seen[i][j] = true
				st = append(st, Point{x: i, y: j})
				t = 0
			} else {
				s += 1
			}
		}
	}

	if t == -1 {
		if s == 0 {
			return 0
		}
		return -1
	}

	for len(st) > 0 {
		p := st[0]
		st = st[1:]
		for i := 0; i < 4; i++ {
			px, py := p.x + dx[i], p.y + dy[i]
			if px >= 0 && px < m && py >= 0 && py < n && !seen[px][py] && grid[px][py] == 1 {
				grid[px][py] = grid[p.x][p.y] + 1
				seen[px][py] = true
				st = append(st, Point{x: px, y: py})
				s--
				if grid[px][py] > t {
					t = grid[px][py]
				}
			}
		}
	}

	if s == 0 {
		return t
	}
	return -1
}


