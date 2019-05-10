package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	yLen := len(grid)
	xLen := len(grid[0])

	marked := map[int]map[int]bool{}
	for i := 0; i < xLen; i++ {
		marked[i] = map[int]bool{}
	}

	var mark func(x, y int)
	mark = func(x, y int) {
		marked[x][y] = true
		if x > 0 && grid[y][x-1] == '1' && !marked[x-1][y] {
			mark(x-1, y)
		}
		if x < xLen-1 && grid[y][x+1] == '1' && !marked[x+1][y] {
			mark(x+1, y)
		}
		if y > 0 && grid[y-1][x] == '1' && !marked[x][y-1] {
			mark(x, y-1)
		}
		if y < yLen-1 && grid[y+1][x] == '1' && !marked[x][y+1] {
			mark(x, y+1)
		}
	}

	var ret int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '1' && !marked[x][y] {
				mark(x, y)
				ret++
			}
		}
	}

	return ret
}
