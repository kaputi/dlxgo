package sudoku

import "fmt"

type matrix struct {
	identifier string
	row        []string
}

type RowForDlx struct {
	RowIdentifier  string
	ColIdentifiers []string
}

/*
[
  ["R1C1N1", "R1C1N2", "R1C1N3", "R1C1N4", "R1C1N5", "R1C1N6", "R1C1N7", "R1C1N8", "R1C1N9"],
]
*/

/*
IDENTIDIERS
position R1C1 R1C2 R1C3...
row constraints  R1N1 R1N2
*/

func GenerateRows2(board [][]int) ([]string, []RowForDlx) {
	rows := []RowForDlx{}
	colsIdentifiers := []string{}
	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			colsIdentifiers = append(colsIdentifiers, fmt.Sprintf("R%dC%d", r, c))
			for n := 1; n <= 9; n++ {
				row := RowForDlx{
					RowIdentifier: fmt.Sprintf("R%dC%dN%d", r, c, n),
					ColIdentifiers: []string{
						fmt.Sprintf("R%d", n-1+(r-1)*9),
						fmt.Sprintf("C%d", n-1+(c-1)*9),
						fmt.Sprintf("B%d", n-1+((r-1)/3*3+(c-1)/3)*9),
					},
				}

				if board[r][c] == 0 {
					row.ColIdentifiers = append(row.ColIdentifiers, fmt.Sprintf("R%dC%d", r, c))
				}

				rows = append(rows, row)
			}
		}
	}

	rowConstraintsIdentifiers := []string{}
	colContraintsIdentifiers := []string{}
	boxContraintsIdentifiers := []string{}

	for i := 0; i < 81; i++ {
		rowConstraintsIdentifiers = append(rowConstraintsIdentifiers, fmt.Sprintf("R%d", i))
		colContraintsIdentifiers = append(colContraintsIdentifiers, fmt.Sprintf("C%d", i))
		boxContraintsIdentifiers = append(boxContraintsIdentifiers, fmt.Sprintf("B%d", i))
	}

	colsIdentifiers = append(colsIdentifiers, rowConstraintsIdentifiers...)
	colsIdentifiers = append(colsIdentifiers, colContraintsIdentifiers...)
	colsIdentifiers = append(colsIdentifiers, boxContraintsIdentifiers...)

	return colsIdentifiers, rows
}

// var fillMatrix = "number"
// var fillMatrix = "row"
// var fillMatrix = "column"

var fillMatrix = "box"

// var fillMatrix = "allNumbers"

func GenerateRows(board [][]int) (filledMatrix []matrix, rows [][]string, identifiers []string) {
	// rows = make([][]string, 729)
	// identifiers = make([]string, 324)

	zeroRow := make([]string, 81)
	for i := 0; i < 81; i++ {
		zeroRow[i] = "."
	}

	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			colIdentifier := fmt.Sprintf("R%dC%d", r, c)
			identifiers = append(identifiers, colIdentifier)
			for n := 1; n <= 9; n++ {
				rowIdentifier := fmt.Sprintf("R%dC%dN%d", r, c, n)
				matrixRow := make([]string, 81)
				copy(matrixRow, zeroRow)

				row := []string{}

				boardNumber := board[r-1][c-1]
				if boardNumber == 0 {
					// if there is no clue, we add nodes for all the numbers
					row = append(row, colIdentifier) // add node in this column for this node

					if fillMatrix == "number" {
						matrixRow[(r-1)*9+c-1] = fmt.Sprint(n)
					}
				} else {
					if boardNumber == n {
						row = append(row, colIdentifier) // add node for specific number
						if fillMatrix == "number" {
							matrixRow[(r-1)*9+c-1] = "X"
						}
					}
				}

				// row constraints
				row = append(row, fmt.Sprintf("R%d", n-1+(r-1)*9))
				// column constraints
				row = append(row, fmt.Sprintf("C%d", n-1+(c-1)*9))
				// box constraints
				row = append(row, fmt.Sprintf("B%d", n-1+((r-1)/3*3+(c-1)/3)*9))

				switch fillMatrix {
				case "row":
					matrixRow[n-1+(r-1)*9] = fmt.Sprint(n)
				case "column":
					matrixRow[n-1+(c-1)*9] = fmt.Sprint(n)
				case "box":
					matrixRow[n-1+((r-1)/3*3+(c-1)/3)*9] = fmt.Sprint(n)
				case "allNumbers":
					matrixRow[(r-1)*9+c-1] = fmt.Sprint(n)
				}

				// TODO: add constraints identifiers

				filledMatrix = append(filledMatrix, matrix{
					identifier: rowIdentifier,
					row:        matrixRow},
				)

				rows = append(rows, row)
			}
		}
	}

	rowConstraintsIdentifiers := []string{}
	colContraintsIdentifiers := []string{}
	boxContraintsIdentifiers := []string{}

	for i := 0; i < 81; i++ {
		rowConstraintsIdentifiers = append(rowConstraintsIdentifiers, fmt.Sprintf("R%d", i))
		colContraintsIdentifiers = append(colContraintsIdentifiers, fmt.Sprintf("C%d", i))
		boxContraintsIdentifiers = append(boxContraintsIdentifiers, fmt.Sprintf("B%d", i))
	}

	identifiers = append(identifiers, rowConstraintsIdentifiers...)
	identifiers = append(identifiers, colContraintsIdentifiers...)
	identifiers = append(identifiers, boxContraintsIdentifiers...)

	// return filledMatrix, rows, identifiers
	return
}
