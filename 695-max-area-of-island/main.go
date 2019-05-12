package main

func maxAreaOfIsland(grid [][]int) int {
	yLength := len(grid)
	if yLength == 0 {
		return 0
	}
	xLength := len(grid[0])
	if xLength == 0 {
		return 0
	}

	searched := map[int]map[int]bool{}
	for y := range grid {
		searched[y] = map[int]bool{}
		for x := range grid[y] {
			searched[y][x] = false
		}
	}

	var measure func(x, y, area int) int
	measure = func(x, y, area int) int {
		if x < 0 || x > xLength-1 || y < 0 || y > yLength-1 || grid[y][x] != 1 || searched[y][x] {
			return area
		}

		area++
		searched[y][x] = true

		area = measure(x+1, y, area)
		area = measure(x-1, y, area)
		area = measure(x, y+1, area)
		area = measure(x, y-1, area)

		return area
	}

	var maxArea int
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 1 {
				area := measure(x, y, 0)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
