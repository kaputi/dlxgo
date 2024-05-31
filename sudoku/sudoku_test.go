package sudoku

import (
	"fmt"
	"testing"

	"github.com/kaputi/dlxgo/newdlx"
)

var board1 = [][]int{
	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	// -------------------
	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	// -------------------
	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	{0, 0, 0, 0, 8, 0, 0, 7, 9},
}

// two solutions
// 	var board1 = [][]int{
//   {2, 9, 5, 7, 4, 3, 8, 6, 1},
// 	{4, 3, 1, 8, 6, 5, 9, 0, 0},
// 	{8, 7, 6, 1, 9, 2, 5, 4, 3},
// 	{3, 8, 7, 4, 5, 9, 2, 1, 6},
// 	{6, 1, 2, 3, 8, 7, 4, 9, 5},
// 	{5, 4, 9, 2, 1, 6, 7, 3, 8},
// 	{7, 6, 3, 5, 2, 4, 1, 8, 9},
// 	{9, 2, 8, 6, 7, 1, 3, 5, 4},
// 	{1, 5, 4, 9, 3, 8, 6, 0, 0},
// }

// EXTRA HARD
// var board1 = [][]int{
// 	{0, 0, 0, 0, 0, 0, 0, 0, 5},
// 	{0, 9, 0, 0, 0, 0, 6, 0, 4},
// 	{0, 7, 0, 0, 0, 0, 2, 0, 0},

// 	{0, 6, 4, 0, 0, 8, 0, 0, 0},
// 	{8, 0, 5, 0, 2, 0, 1, 0, 0},
// 	{2, 0, 0, 0, 7, 0, 0, 0, 3},

// 	{0, 2, 0, 0, 0, 7, 0, 3, 0},
// 	{0, 0, 0, 0, 0, 6, 0, 0, 1},
// 	{0, 0, 1, 9, 0, 0, 0, 4, 0},
// }

var solution1 = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

func TestSudoku(t *testing.T) {

	printMatrix := false
	printRows := true
	printIdentifiers := false

	matrix, rows, identifiers := GenerateRows(board1)

	if printMatrix {
		for _, row := range matrix {
			name := row.identifier
			values := row.row
			fmt.Printf("%v->%v\n", name, values)
		}
	}

	if printRows {
		for i, row := range rows {
			number := (i % 9) + 1
			fmt.Printf("%d -> %v\n", number, row)
			if number == 9 {
				println("----------")
			}
		}
	}

	if printIdentifiers {
		fmt.Println(identifiers)
	}

	dlx := newdlx.NewDlxMatrix(identifiers)

	for i, row := range rows {
		idetifier := fmt.Sprintf("%d", i%9+1)
		dlx.AddConstraintRow(idetifier, row)
	}

	// dlx.SolveOne()
	dlx.SolveAll()

	// fmt.Println(dlx.GetSolution())

	// solution, board := dlx.SOLSOL()
	// fmt.Println(solution)

	// _, boards := dlx.SOLSOL()
	// for i, board := range boards {
	// 	fmt.Println("Solution: ", i+1)
	// 	for _, row := range board {
	// 		fmt.Println(row)
	// 	}
	// }

}
