package main

import (
	"math"
	"sort"
)

func sampleStats(count []int) []float64 {
	min := math.MaxFloat64
	var countMax, sampleCount int
	var max, mean, median, mode, sum float64
	for i := 0; i < len(count); i++ {
		if count[i] != 0 && float64(i) < min {
			min = float64(i)
		}
		if count[i] != 0 && float64(i) > max {
			max = float64(i)
		}
		if count[i] > countMax {
			countMax = count[i]
			mode = float64(i)
		}
		sum += float64(i) * float64(count[i])
		sampleCount += count[i]
	}

	if sampleCount == 0 {
		mean = 0
		min = 0
	} else {
		mean = sum / float64(sampleCount)
	}

	if sampleCount == 0 {
		median = 0
	} else if sampleCount%2 == 0 {
		var p, c, a, b int
		for i := 0; i < len(count); i++ {
			c += count[i]
			if p < sampleCount/2 && sampleCount/2 <= c {
				a = i
			}
			if p < sampleCount/2+1 && sampleCount/2+1 <= c {
				b = i
			}
			p = c
		}

		median = float64(a+b) / float64(2)
	} else {
		var p, c, a int
		for i := 0; i < len(count); i++ {
			c += count[i]
			if p < sampleCount/2 && sampleCount/2 <= c {
				a = i
			}
			p = c
		}

		median = float64(a)
	}

	return []float64{min, max, mean, median, mode}
}

func carPooling(trips [][]int, capacity int) bool {
	var eventQueue [][]int
	for i := 0; i < len(trips); i++ {
		eventQueue = append(eventQueue, []int{trips[i][0], trips[i][1]})
		eventQueue = append(eventQueue, []int{-trips[i][0], trips[i][2]})
	}

	sort.Slice(eventQueue, func(i, j int) bool {
		if eventQueue[i][1] == eventQueue[j][1] {
			return eventQueue[i][0] < eventQueue[j][0]
		}
		return eventQueue[i][1] < eventQueue[j][1]
	})

	var current int
	for i := 0; i < len(eventQueue); i++ {
		current += eventQueue[i][0]
		if current > capacity {
			return false
		}
	}

	return true
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}

func (this *MountainArray) length() int {
	return len(this.arr)
}

type MountainArray struct {
	arr []int
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	min := 0
	arrLength := mountainArr.length()
	max := arrLength - 1
	var i int
	for {
		mid := min + (max-min)/2
		midVal := mountainArr.get(mid)

		if mid > 0 {
			a := mountainArr.get(mid - 1)
			if a > midVal {
				max = mid - 1
				continue
			} else if a < midVal {
				if mid < arrLength-1 {
					b := mountainArr.get(mid + 1)
					if b > midVal {
						min = mid + 1
						continue
					} else if b < midVal {
						i = mid
						break
					}
				} else {
					i = mid
					break
				}
			}
		} else {
			i = 1
			break
		}
	}

	min = 0
	max = i
	ret := -1
	for {
		if max < min {
			break
		} else {
			mid := min + (max-min)/2
			val := mountainArr.get(mid)
			if val > target {
				max = mid - 1
				continue
			} else if val < target {
				min = mid + 1
				continue
			} else {
				ret = mid
				break
			}
		}
	}

	if ret == -1 {
		min = i
		max = arrLength - 1
		for {
			if max < min {
				break
			} else {
				mid := min + (max-min)/2
				val := mountainArr.get(mid)
				if val < target {
					max = mid - 1
					continue
				} else if val > target {
					min = mid + 1
					continue
				} else {
					ret = mid
					break
				}
			}
		}
	}

	return ret
}

func parseBrace(expression string) (map[string]bool, int) {
	ret := map[string]bool{}
	current := map[string]bool{}
	var i int
	for {
		if i >= len(expression) {
			for k := range current {
				ret[k] = true
			}
			return ret, i
		}
		if expression[i] == '{' {
			list, index := parseBrace(expression[i+1:])
			if len(current) == 0 {
				for k := range list {
					current[k] = true
				}
			} else {
				nextCurrent := map[string]bool{}
				for k1 := range current {
					for k2 := range list {
						nextCurrent[k1+k2] = true
					}
				}
				current = nextCurrent
			}
			i += index + 2
		} else if expression[i] == ',' {
			for k := range current {
				ret[k] = true
			}
			current = map[string]bool{}
			i++
		} else if expression[i] == '}' {
			for k := range current {
				ret[k] = true
			}
			return ret, i
		} else {
			if len(current) == 0 {
				current[string(expression[i])] = true
			} else {
				nextCurrent := map[string]bool{}
				for k := range current {
					nextCurrent[k+string(expression[i])] = true
				}
				current = nextCurrent
			}
			i++
		}
	}
}

func braceExpansionII(expression string) []string {
	ret, _ := parseBrace(expression)
	var keys []string
	for k := range ret {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
