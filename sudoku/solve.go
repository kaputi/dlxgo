package sudoku

func solveOne(board *boardT) []string {
	dlx := generateDlx(board)
	solution := dlx.SolveOne()
	return solution
}

func solveAll(board *boardT) [][]string {
	dlx := generateDlx(board)
	solutions := dlx.SolveAll()
	return solutions
}
