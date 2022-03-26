package main

var (
	dx = []int{0,1,0,-1}
	dy = []int{1,0,-1,0}
)

func maxAreaOfIsland(grid [][]int) int {
	maxInt := 0
	for m := 0; m < len(grid); m++ {
		for n := 0; n < len(grid[0]); n++ {
			if grid[m][n] == 1 { // 发现了土地开始进行深度搜索
				t := bfs(m, n, grid)
				if t > maxInt {
					maxInt = t
				}
			}
		}
	}
	return maxInt
}

/**
深度优先搜索
 */
func dfs(m, n int, grid [][]int) int {
	grid[m][n] = 0
	st, i, num := &Data{}, 0, 1
	for i < 4 {
		x, y := m + dx[i], n + dy[i]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
			st.pushStack(m, n, i)
			m, n, i = x, y, 0
			grid[x][y] = 0
			num++
			continue
		}
		for i == 3 && !st.isStackEmpty() {
			m, n , i = st.popStack()
		}
		i++
	}
	return num
}

/**
广度优先搜索
 */
func bfs(m, n int, grid [][]int) int {
	st, i, num := &Data{}, 0, 0
	for i < 4 {
		x, y := m + dx[i], n + dy[i]
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
			st.pushQueue(x, y, 0)
		}
		if i == 3 {
			grid[m][n] = 0
			num += 1
			for !st.isQueueEmpty() {
				m, n, _ = st.popQueue()
				if grid[m][n] == 1 {
					i = -1
					break
				}
			}
		}
		i++
	}
	return num
}



type Data struct {
	xData []int
	yData []int
	iData []int
}

func (t *Data) pushStack(x, y, i int) {
	t.xData = append(t.xData, x)
	t.yData = append(t.yData, y)
	t.iData = append(t.iData, i)
}

func (t *Data) popStack() (int, int, int) {
	length := len(t.iData)
	if length == 0 {
		panic(length)
	}
	a, b, c := t.xData[length - 1], t.yData[length - 1], t.iData[length - 1]
	t.xData = t.xData[0:length - 1]
	t.yData = t.yData[0:length - 1]
	t.iData = t.iData[0:length - 1]
	return a, b, c
}

func (t *Data) isStackEmpty() bool {
	return len(t.iData) == 0
}





func (t *Data) pushQueue(x, y, i int) {
	t.xData = append(t.xData, x)
	t.yData = append(t.yData, y)
	t.iData = append(t.iData, i)
}

func (t *Data) popQueue() (int, int, int) {
	length := len(t.iData)
	if length == 0 {
		panic(length)
	}
	a, b, c := t.xData[0], t.yData[0], t.iData[0]
	t.xData = t.xData[1:length]
	t.yData = t.yData[1:length]
	t.iData = t.iData[1:length]
	return a, b, c
}

func (t *Data) isQueueEmpty() bool {
	return len(t.iData) == 0
}




func main() {
	grid := [][]int{
		{1,1,0,0,0},
		{1,1,0,0,0},
		{0,0,0,1,1},
		{0,0,0,1,1}}

	println(maxAreaOfIsland(grid))
}
