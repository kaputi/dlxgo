package ssudoku

import (
	"fmt"
	"math/rand"
)

func generateSolvedBoard() boardT {
	board := boardT{}
	// fill randomly with numbers from 1 to 9
	// for i := 1; i <= 17; i++ {
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

	bruteForce(&board)

	// dlx := generateDlx(board)
	// solutions := dlx.SolveAll()
	// if dlx != nil {
	// 	fmt.Println("DLX generated")
	// 	fmt.Println(solutions)
	// }

	// solution := solveOne(board)
	// solution := []string{}

	// fmt.Println(solution)

	fmt.Println()

	return board
}
