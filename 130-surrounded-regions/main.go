package main

var (
	xByte = []byte("X")[0]
	oByte = []byte("O")[0]
)

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	var answerBoard [][]byte

	yLen := len(board)
	xLen := len(board[0])
	for i := 0; i < yLen; i++ {
		var line []byte
		for j := 0; j < xLen; j++ {
			line = append(line, xByte)
		}
		answerBoard = append(answerBoard, line)
	}

	for y, line := range board {
		for x, b := range line {
			if b == xByte {
				continue
			}

			if (y == 0 || x == 0 || y == yLen-1 || x == xLen-1) && answerBoard[y][x] != oByte {
				answerBoard[y][x] = oByte
				markRecursively(board, answerBoard, y, x, yLen, xLen, map[int]struct{}{})
			}
		}
	}

	for y, line := range board {
		for x := range line {
			if answerBoard[y][x] == xByte {
				board[y][x] = xByte
			}
		}
	}
}

func markRecursively(board [][]byte, answerBoard [][]byte, y int, x int, yLen int, xLen int, c map[int]struct{}) {
	if _, ok := c[y*xLen+x]; ok {
		return
	}
	c[y*xLen+x] = struct{}{}

	// up
	if y > 0 && board[y-1][x] == oByte {
		answerBoard[y-1][x] = oByte
		markRecursively(board, answerBoard, y-1, x, yLen, xLen, c)
	}
	// left
	if x > 0 && board[y][x-1] == oByte {
		answerBoard[y][x-1] = oByte
		markRecursively(board, answerBoard, y, x-1, yLen, xLen, c)
	}
	// right
	if x+1 < xLen && board[y][x+1] == oByte {
		answerBoard[y][x+1] = oByte
		markRecursively(board, answerBoard, y, x+1, yLen, xLen, c)
	}
	// down
	if y+1 < yLen && board[y+1][x] == oByte {
		answerBoard[y+1][x] = oByte
		markRecursively(board, answerBoard, y+1, x, yLen, xLen, c)
	}
}
