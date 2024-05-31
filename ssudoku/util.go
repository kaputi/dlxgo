package ssudoku

import "math/rand"

func getNextEmpty(board *boardT) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j].Value == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func getFilledCoords(board *boardT) [][2]int {
	coords := [][2]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j].Value != 0 {
				coords = append(coords, [2]int{i, j})
			}
		}
	}
	return coords
}

func isValid(board *boardT, row, col, value int) bool {
	if value == 0 {
		return true
	}

	for i := 0; i < 9; i++ {
		if board[row][i].Value == value && i != col {
			return false
		}
		if board[i][col].Value == value && i != row {
			return false
		}
	}

	firstRowInBox := row - row%3
	firstColInBox := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checkRow := firstRowInBox + i
			checkCol := firstColInBox + j
			if checkRow == row && checkCol == col {
				continue
			}
			if board[checkRow][checkCol].Value == value {
				return false
			}
		}
	}

	return true
}

func boardIsFull(board *boardT) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j].Value == 0 {
				return false
			}
		}
	}
	return true
}

var directions = [][2]int{
	{1, 1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

func getDiagonalCoords(coord [2]int, count int) [][2]int {
	randomDirection := directions[rand.Intn(len(directions))]

	diagonals := [][2]int{}

	for i := 1; i <= count; i++ {
		row := coord[0] + randomDirection[0]*i
		col := coord[1] + randomDirection[1]*i
		if row >= 0 && row <= 8 && col >= 0 && col <= 8 {
			diagonals = append(diagonals, [2]int{row, col})
		}
		row = coord[0] - randomDirection[0]*-i
		col = coord[1] - randomDirection[1]*-i
		if row >= 0 && row <= 8 && col >= 0 && col <= 8 {
			diagonals = append(diagonals, [2]int{row, col})
		}
	}

	return diagonals
}
