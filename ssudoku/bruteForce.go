package ssudoku

import (
	"time"
)

func bruteForce(board *boardT) bool {
	emptySpaces := [][]int{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col].Value == 0 {
				emptySpaces = append(emptySpaces, []int{row, col})
			}
		}
	}

	start := time.Now()

	return bruteForceHelper(board, &emptySpaces, 0, &start)
}

func bruteForceHelper(board *boardT, emtpySpaces *[][]int, index int, start *time.Time) bool {
	if index >= len(*emtpySpaces) {
		return true
	}

	row := (*emtpySpaces)[index][0]
	col := (*emtpySpaces)[index][1]

	for value := 1; value <= 9; value++ {
		if time.Since(*start) > 1*time.Second {
			return false
		}
		if !isValid(board, row, col, value) {
			continue
		}

		board[row][col].Value = value

		if bruteForceHelper(board, emtpySpaces, index+1, start) {
			return true
		}

		board[row][col].Value = 0
	}
	return false
}
