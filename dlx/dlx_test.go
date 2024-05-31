package dlx

import (
	"fmt"
	"testing"
)

type dlxTest struct {
	identifiers []string
	rows        [][]string
	expects     [][][]string
}

func TestDdlx(t *testing.T) {
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
			expects: [][][]string{
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
			expects: [][][]string{{{"1", "4"}, {"3", "5", "6"}, {"2", "7"}}},
		},
		// 2
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows:        [][]string{{"2", "3"}},
			expects:     [][][]string{},
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
			expects: [][][]string{{{"2", "4", "5"}, {"0", "3"}, {"1", "6"}}},
		},
		// // 4
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows: [][]string{
				{"0", "1"},
				{"0", "2"},
				{"1", "2"},
			},
			expects: [][][]string{},
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
			expects: [][][]string{
				{{"0", "1", "2"}, {"3"}},
				{{"0", "2"}, {"1"}, {"3"}},
			},
		},
	}

	for i, test := range tests {
		dlx := NewDlx(test.identifiers)
		for i, row := range test.rows {
			dlx.AddConstraintRow(fmt.Sprint(i), row)
		}

		solutionRows := dlx.SolveAll()
		solutionCols := dlx.GetColNamesFromSolutions()

		fmt.Printf("test %v solutions ===================\n", i)
		fmt.Printf("solution rows: %v\n", solutionRows)
		fmt.Printf("solution names: %v\n", solutionCols)
		fmt.Printf("expected: %v\n", test.expects)

		// for _, solution := range solutions {
		// fmt.Printf("solution: %v\n", solution)
		// fmt.Printf("expected: %v\n", test.expects)
		// }
	}
}
