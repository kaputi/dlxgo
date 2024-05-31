package sudoku

import (
	"math/rand"
)

func scrambleSlice[T any](slice *[]T) {
	for i := range *slice {
		j := rand.Intn(i + 1)
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

func generateSolvedBoard() boardT {
	board := boardT{}
	// fill randomly with numbers from 1 to 9
	for i := 1; i <= 9; i++ {
		row := rand.Intn(9)
		col := rand.Intn(9)
		for board[row][col].Value != 0 && isValid(&board, row, col, i) {
			row = rand.Intn(9)
			col = rand.Intn(9)
		}

		// board[row][col].Value = i%9 + 1
		board[row][col].Value = i
	}

	if !bruteForce(&board) {
		board = generateSolvedBoard()
	}

	scrambled := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	scrambleSlice(&scrambled)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			board[row][col].Value = scrambled[board[row][col].Value-1]
		}
	}

	return board
}

func generate(dificulty int) boardT {
	dificulty = max(0, min(6, dificulty))

	board := generateSolvedBoard()
	iterations := 0
	removeNumbers(&board, dificulty, &iterations)

	for rowI, row := range board {
		for colI, cell := range row {
			if cell.Value != 0 {
				board[rowI][colI].Fixed = true
			}
		}
	}

	return board
}

func removeNumbers(board *boardT, dificulty int, iterations *int) {
	dificulty = max(0, min(6, dificulty))
	clues := 30 - dificulty

	filledCoords := getFilledCoords(board)
	filledLen := len(filledCoords)

	if filledLen <= clues {
		return
	}

	coord := filledCoords[rand.Intn(filledLen)]
	if !removeCoordIfPosible(board, coord) {
		*iterations++
		if *iterations > 20 {
			return
		}

		removeNumbers(board, dificulty, iterations)
	} else {
		filledLen--
	}

	if filledLen <= clues {
		return
	}

	diagonalCount := 2
	if filledLen <= 60 {
		diagonalCount = 1
	}

	coordDiagonals := getDiagonalCoords(coord, diagonalCount)

	for _, coord := range coordDiagonals {
		if removeCoordIfPosible(board, coord) {
			filledLen--
			if filledLen <= clues {
				return
			}
		}
	}

	removeNumbers(board, dificulty, iterations)
}

func removeCoordIfPosible(board *boardT, coord [2]int) bool {
	valBackup := board[coord[0]][coord[1]].Value
	(*board)[coord[0]][coord[1]].Value = 0

	dlx := generateDlx(board)
	solutions := dlx.SolveAll()
	if len(solutions) != 1 {
		(*board)[coord[0]][coord[1]].Value = valBackup
		return false
	}

	return true
}
