package section746

func minCostClimbingStairs(cost []int) int {
	sum, n0, n1 := 0, 0, 0
	for i, n := 0, len(cost); i <= n; i++ {
		if i == 0 || i == 1 {
			continue
		}
		sum = min(n0 + cost[i - 2], n1 + cost[i -1])
		n0 = n1
		n1 = sum
	}
	return sum
}

func min(a, b int) int {
	if b <= a {
		return b
	}
	return a
}

func main() {
	println(minCostClimbingStairs([]int{1,100,1,1,1,100,1,1,100,1}))
}