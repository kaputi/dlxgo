package sudoku

type cell struct {
	Value          int
	Fixed          bool
	PlacementError bool
	SolutionError  bool
}

type boardT [9][9]cell

var PlacedLayer = "placed"
var Notes1Layer = "notes1"
var Notes2Layer = "notes2"

type Sudoku struct {
	solved boardT
	placed boardT
	notes1 boardT
	notes2 boardT
	layer  string
}

func NewSudoku() *Sudoku {
	return &Sudoku{}
}

func (s *Sudoku) SetLayer(layer string) {
	s.layer = layer
}

func boardFromMatrix(matrix [][]int) boardT {
	board := boardT{}
	for rowI, row := range matrix {
		for colI, val := range row {
			board[rowI][colI].Value = val
		}
	}
	return board
}
