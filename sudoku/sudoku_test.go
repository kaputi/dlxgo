package sudoku

import (
	"fmt"
	"strconv"
	"github.com/kaputi/dlxgo/util"
	"testing"
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

var multipleSolutions = [][]int{
	{2, 9, 5, 7, 4, 3, 8, 6, 1},
	{4, 3, 1, 8, 6, 5, 9, 0, 0},
	{8, 7, 6, 1, 9, 2, 5, 4, 3},
	{3, 8, 7, 4, 5, 9, 2, 1, 6},
	{6, 1, 2, 3, 8, 7, 4, 9, 5},
	{5, 4, 9, 2, 1, 6, 7, 3, 8},
	{7, 6, 3, 5, 2, 4, 1, 8, 9},
	{9, 2, 8, 6, 7, 1, 3, 5, 4},
	{1, 5, 4, 9, 3, 8, 6, 0, 0},
}

func TestSolveWithDlx(t *testing.T) {

	board := boardFromMatrix(board1)

	dlx := generateDlx(&board)
	solution := dlx.SolveAll()
	for _, s := range solution {
		// fmt.Printf("Solution: %v\n", s)
		solutionMtx := make([][]int, 9)
		for i := range solutionMtx {
			solutionMtx[i] = make([]int, 9)
		}

		errs := util.NewErrs()
		for _, cell := range s {
			rowStr, err := strconv.Atoi(fmt.Sprintf("%c", cell[1]))
			errs.Add(err)
			colStr, err := strconv.Atoi(fmt.Sprintf("%c", cell[3]))
			errs.Add(err)
			numStr, err := strconv.Atoi(fmt.Sprintf("%c", cell[5]))
			errs.Add(err)
			if errs.Has() {
				errs.Print()
				t.Fail()
			}
			solutionMtx[rowStr-1][colStr-1] = numStr
		}

		for _, row := range solutionMtx {
			fmt.Println(row)
		}
	}
}

func TestGenerateSolved(t *testing.T) {

	board := generateSolvedBoard()
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%d ", cell.Value)
		}
		fmt.Println()
	}
}

func TestGenerate(t *testing.T) {

	board := generate(4)
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%d ", cell.Value)
		}
		fmt.Println()
	}
}
