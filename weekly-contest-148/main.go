package main

func movesToMakeZigzag(nums []int) int {
	even := 0
	odd := 0
	evenPrev := nums[0]
	oddPrev := nums[0]
	for i := 1; i < len(nums); i++ {
		evenN := nums[i]
		oddN := nums[i]
		if i%2 == 0 {
			if evenN <= evenPrev {
				even += evenPrev - evenN + 1
			}
			if oddN >= oddPrev {
				odd += oddN - oddPrev + 1
				oddN = oddPrev - 1
			}
		} else {
			if evenN >= evenPrev {
				even += evenN - evenPrev + 1
				evenN = evenPrev - 1
			}
			if oddN <= oddPrev {
				odd += oddPrev - oddN + 1
			}
		}
		evenPrev = evenN
		oddPrev = oddN
	}

	if even < odd {
		return even
	}
	return odd
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	xCount := 0
	parent := 0
	right := 0
	left := 0
	var r func(node *TreeNode, prevVal int, count bool)
	r = func(node *TreeNode, prevVal int, count bool) {
		if node.Val == x {
			count = true
			parent = prevVal
			if node.Left != nil {
				left = node.Left.Val
			}
			if node.Right != nil {
				right = node.Right.Val
			}
		}
		if count {
			xCount++
		}
		if node.Left != nil {
			r(node.Left, node.Val, count)
		}
		if node.Right != nil {
			r(node.Right, node.Val, count)
		}
	}
	r(root, 0, false)

	var total int
	var rCount func(node *TreeNode, target int, count bool)
	rCount = func(node *TreeNode, target int, count bool) {
		if node.Val == target {
			count = true
		}
		if count {
			total++
		}
		if node.Left != nil {
			rCount(node.Left, target, count)
		}
		if node.Right != nil {
			rCount(node.Right, target, count)
		}
	}

	if parent != 0 {
		total = 0
		rCount(root, parent, false)
		if n-xCount > xCount {
			return true
		}
	}
	if left != 0 {
		total = 0
		rCount(root, left, false)
		if total > n-total {
			return true
		}
	}
	if right != 0 {
		total = 0
		rCount(root, right, false)
		if total > n-total {
			return true
		}
	}
	return false
}

type SnapshotArray struct {
	current         []int
	snapshots       [][]int
	snapshotHistory []int
	snapshotIndex   int
	history         [][]int
	historyIndex    int
}

func Constructor(length int) SnapshotArray {
	current := make([]int, length)
	snapshots := make([][]int, 500)
	snapshots[0] = make([]int, len(current))
	copy(snapshots[0], current)
	snapshotHistory := make([]int, 50000)
	history := make([][]int, 50000)

	return SnapshotArray{
		current:         current,
		snapshots:       snapshots,
		snapshotHistory: snapshotHistory,
		snapshotIndex:   0,
		history:         history,
		historyIndex:    0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.current[index] = val
	this.history[this.historyIndex] = []int{index, val}
	this.historyIndex++

	if this.historyIndex%100 == 0 {
		index := this.historyIndex / 100
		this.snapshots[index] = make([]int, len(this.current))
		copy(this.snapshots[index], this.current)
	}
}

func (this *SnapshotArray) Snap() int {
	snapId := this.snapshotIndex
	this.snapshotHistory[snapId] = this.historyIndex
	this.snapshotIndex++
	return snapId
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	h := this.snapshotHistory[snap_id]
	id := h / 100
	m := h % 100
	ret := this.snapshots[id][index]

	for i := 0; i < m; i++ {
		if this.history[id*100+i][0] == index {
			ret = this.history[id*100+i][1]
		}
	}

	return ret
}

func longestDecomposition(text string) int {
	if len(text) == 0 {
		return 0
	}
	for i := 1; i <= len(text)/2; i++ {
		if text[0:i] == text[len(text)-i:] {
			return longestDecomposition(text[i:len(text)-i]) + 2
		}
	}
	return 1
}
