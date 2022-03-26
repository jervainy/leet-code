package main

var (
	dx = []int{0,1,0,-1}
	dy = []int{1,0,-1,0}
)


func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	cColor := image[sr][sc]
	if cColor != newColor {
		dfs(image, sr, sc, newColor, cColor)
	}
	return image
}

func dfs(image [][]int, r int, c int, newColor, oColor int) {
	if image[r][c] == oColor {
		image[r][c] = newColor
		for i := 0; i < 4; i++ {
			x, y := r + dx[i], c + dy[i]
			if x >= 0 && x < len(image) && y >= 0 && y < len(image[0]) {
				dfs(image, x, y, newColor, oColor)
			}
		}
	}
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	floodFill(image, 1, 1, 2)
	println(image)
}

