package newdlx

import (
	"fmt"
	"testing"
)

type dlxTest struct {
	identifiers []string
	rows        [][]string
	solutions   [][][]string
}

func TestNewDlx(t *testing.T) {

	tests := []dlxTest{
		// 0
		{
			identifiers: []string{"A", "B", "C", "D"},
			rows: [][]string{
				{"A", "B"},
				{"C"},
				{"C", "D"},
				{"D"},
			},
			solutions: [][][]string{
				{{"A", "B"}, {"C"}, {"D"}},
				{{"A", "B"}, {"C", "D"}},
			},
		},
		// 1
		{
			identifiers: []string{"1", "2", "3", "4", "5", "6", "7"},
			rows: [][]string{
				{"1", "4", "7"},
				{"1", "4"},
				{"4", "5", "7"},
				{"3", "5", "6"},
				{"2", "3", "6", "7"},
				{"2", "7"},
			},
			solutions: [][][]string{{{"1", "4"}, {"3", "5", "6"}, {"2", "7"}}},
		},
		// 2
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows:        [][]string{{"2", "3"}},
			solutions:   [][][]string{},
		},
		// // 3
		{
			identifiers: []string{"0", "1", "2", "3", "4", "5", "6"},
			rows: [][]string{
				{"2", "4", "5"},
				{"0", "3", "6"},
				{"1", "2", "5"},
				{"0", "3"},
				{"1", "6"},
				{"3", "4", "6"},
			},
			solutions: [][][]string{{{"2", "4", "5"}, {"0", "3"}, {"1", "6"}}},
		},
		// // 4
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows: [][]string{
				{"0", "1"},
				{"0", "2"},
				{"1", "2"},
			},
			solutions: [][][]string{},
		},
		// // 5
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows: [][]string{
				{"0", "1", "2"},
				{"0", "2"},
				{"1"},
				{"3"},
			},
			// 6
			solutions: [][][]string{
				{{"0", "1", "2"}, {"3"}},
				{{"0", "2"}, {"1"}, {"3"}},
			},
		},
	}

	for _, test := range tests {
		dlx := NewDlxMatrix(test.identifiers)
		for i, row := range test.rows {
			dlx.AddConstraintRow(fmt.Sprint(i), row)
		}
		dlx.debug = true

		dlx.SolveAll()

		fmt.Println()
		fmt.Println("test solution: ", dlx.solutions)
		fmt.Println("expected solution: ", test.solutions)
		fmt.Println()
	}
}
