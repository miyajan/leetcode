package main

func maxArea(height []int) int {
	var maxArea int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			var h int
			if height[i] < height[j] {
				h = height[i]
			} else {
				h = height[j]
			}
			area := w * h
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
