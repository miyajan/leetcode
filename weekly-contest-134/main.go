package main

import (
	"math"
)

func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}

	var min int
	if b-a == 2 || c-b == 2 {
		min = 1
	} else {
		if b-a > 1 {
			min++
		}
		if c-b > 1 {
			min++
		}
	}

	max := (b - a - 1) + (c - b - 1)

	return []int{min, max}
}

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	var colored [][]int
	checked := make(map[int]map[int]bool)

	colored = getColoredRecursively(grid, r0, c0, grid[r0][c0], colored, checked)
	for _, g := range colored {
		r := g[0]
		c := g[1]
		grid[r][c] = color
	}

	return grid
}

func getColoredRecursively(grid [][]int, r int, c int, baseColor int, colored [][]int, checked map[int]map[int]bool) [][]int {
	if _, ok := checked[r]; !ok {
		checked[r] = make(map[int]bool)
	}
	checked[r][c] = true
	if grid[r][c] != baseColor {
		return colored
	}

	isBorder := false
	if r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[0])-1 {
		isBorder = true
	}
	if r > 0 {
		if grid[r-1][c] != grid[r][c] {
			isBorder = true
		}
		if !checked[r-1][c] {
			colored = getColoredRecursively(grid, r-1, c, baseColor, colored, checked)
		}
	}
	if r < len(grid)-1 {
		if grid[r+1][c] != grid[r][c] {
			isBorder = true
		}
		if !checked[r+1][c] {
			colored = getColoredRecursively(grid, r+1, c, baseColor, colored, checked)
		}
	}
	if c > 0 {
		if grid[r][c-1] != grid[r][c] {
			isBorder = true
		}
		if !checked[r][c-1] {
			colored = getColoredRecursively(grid, r, c-1, baseColor, colored, checked)
		}
	}
	if c < len(grid[0])-1 {
		if grid[r][c+1] != grid[r][c] {
			isBorder = true
		}
		if !checked[r][c+1] {
			colored = getColoredRecursively(grid, r, c+1, baseColor, colored, checked)
		}
	}

	if isBorder {
		colored = append(colored, []int{r, c})
	}

	return colored
}

func maxUncrossedLines(A []int, B []int) int {
	dp := make(map[int]map[int]int)
	for i := 0; i <= len(A); i++ {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 0
	dp[0][1] = 0
	dp[1][0] = 0
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i-1][j])))
			}
		}
	}
	return dp[len(A)][len(B)]
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	blockedMap := make(map[int]map[int]bool)
	for i := 0; i < len(blocked); i++ {
		x := blocked[i][0]
		y := blocked[i][1]
		if _, ok := blockedMap[x]; !ok {
			blockedMap[x] = make(map[int]bool)
		}
		blockedMap[x][y] = true
	}

	return isEscapePossibleRecursibly(blockedMap, source[0], source[1], target[0], target[1], 0, make(map[int]map[int]bool)) &&
		isEscapePossibleRecursibly(blockedMap, target[0], target[1], source[0], source[1], 0, make(map[int]map[int]bool))
}

func isEscapePossibleRecursibly(blocked map[int]map[int]bool, sourceX int, sourceY int, targetX int, targetY int, step int, checked map[int]map[int]bool) bool {
	if checked[sourceX][sourceY] ||
		sourceX < 0 ||
		sourceX >= 1000000 ||
		sourceY < 0 ||
		sourceY >= 1000000 ||
		blocked[sourceX][sourceY] {
		return false
	}

	if _, ok := checked[sourceX]; !ok {
		checked[sourceX] = make(map[int]bool)
	}
	checked[sourceX][sourceY] = true

	if sourceX == targetX && sourceY == targetY {
		return true
	}

	if step > 20000 {
		return true
	}

	step++
	return isEscapePossibleRecursibly(blocked, sourceX+1, sourceY, targetX, targetY, step, checked) ||
		isEscapePossibleRecursibly(blocked, sourceX, sourceY+1, targetX, targetY, step, checked) ||
		isEscapePossibleRecursibly(blocked, sourceX-1, sourceY, targetX, targetY, step, checked) ||
		isEscapePossibleRecursibly(blocked, sourceX, sourceY-1, targetX, targetY, step, checked)
}
