package sudoku

import (
	"fmt"

	"github.com/kaputi/dlxgo/dlx"
)

type sudokuDlxRow struct {
	name        string
	columnNames []string
}

func generateDlx(board *boardT) *dlx.Dlx {
	rows := []sudokuDlxRow{}
	colNames := []string{}

	for r := 1; r <= 9; r++ {
		for c := 1; c <= 9; c++ {
			colNames = append(colNames, fmt.Sprintf("R%dC%d", r, c))
			for n := 1; n <= 9; n++ {
				row := sudokuDlxRow{
					name: fmt.Sprintf("R%dC%dN%d", r, c, n),
					columnNames: []string{
						// 	fmt.Sprintf("R%dC%d", r, c),                   // number constarint
						fmt.Sprintf("R%d", n-1+(r-1)*9),               // row constraint
						fmt.Sprintf("C%d", n-1+(c-1)*9),               // column constraint
						fmt.Sprintf("B%d", n-1+((r-1)/3*3+(c-1)/3)*9), // box constraint
					},
				}
				if board[r-1][c-1].Value == 0 || board[r-1][c-1].Value == n {
					row.columnNames = append(row.columnNames, fmt.Sprintf("R%dC%d", r, c))
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

	colNames = append(colNames, rowConstraintsIdentifiers...)
	colNames = append(colNames, colContraintsIdentifiers...)
	colNames = append(colNames, boxContraintsIdentifiers...)

	dlx := dlx.NewDlx(colNames)

	for _, row := range rows {
		// fmt.Printf("name: %v, columns: %v\n", row.name, row.columnNames)
		dlx.AddConstraintRow(row.name, row.columnNames)
	}

	return dlx
}
