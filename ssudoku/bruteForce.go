package ssudoku

import "fmt"

func bruteForce(board *boardT) bool {
	emptySpaces := [][]int{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col].Value == 0 {
				emptySpaces = append(emptySpaces, []int{row, col})
			}
		}
	}

	return bruteForceHelper(board, &emptySpaces, 0)
}

func bruteForceHelper(board *boardT, emtpySpaces *[][]int, index int) bool {
	if index >= len(*emtpySpaces) {
		return true
	}

	row := (*emtpySpaces)[index][0]
	col := (*emtpySpaces)[index][1]

	for value := 1; value <= 9; value++ {
		if !isValid(board, row, col, value) {
			continue
		}

		board[row][col].Value = value

		// fmt.Print("\033[H\033[2J")

    fmt.Println("BRUTE FORCE")
		for _, row := range *board {
			for _, cell := range row {
				fmt.Printf("%d ", cell.Value)
			}
			fmt.Println()
		}

		if bruteForceHelper(board, emtpySpaces, index+1) {
			return true
		}

		board[row][col].Value = 0
	}
	return false
}
