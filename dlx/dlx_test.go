package dlx

import (
	"fmt"
	"testing"
)

type dlxTest struct {
	identifiers []string
	rows        [][]string
	solutions   [][]int // rows included in solutions
}

func TestDlx(t *testing.T) {

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
			solutions: [][]int{{0, 2}, {0, 1, 3}},
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
			solutions: [][]int{{1, 3, 5}},
		},
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows:        [][]string{{"2", "3"}},
			solutions:   [][]int{},
		},
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
			solutions: [][]int{{0, 3, 4}},
		},
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows: [][]string{
				{"0", "1"},
				{"0", "2"},
				{"1", "2"},
			},
			solutions: [][]int{},
		},
		{
			identifiers: []string{"0", "1", "2", "3"},
			rows: [][]string{
				{"0", "1", "2"},
				{"0", "2"},
				{"1"},
				{"3"},
			},
			solutions: [][]int{
				{0, 3},
				{1, 2, 3},
			},
		},
	}

	for _, test := range tests {
		dlx := NewDlx(test.identifiers)
		for _, row := range test.rows {
			dlx.AddConstraintRow(row)
		}
		dlx.Solve2()

		// root := dlx.root
		// curr := root.right
		// for curr != root {
		// 	fmt.Print(curr.identifier)
		//     fmt.Printf(" colSize: %d", curr.colSize)
		// 	fmt.Print(" -> ")
		// 	curr = curr.right
		// }
		// fmt.Println()

		fmt.Println("---------- test------------")
		// for _, solution := range dlx.solutions {
		// 	fmt.Println("solution: ")
		// 	for _, row := range solution {
		// 		rowHead, ok := dlx.rowHeads[row]
		// 		if !ok {
		// 			fmt.Println("row not found")
		// 		} else {
		// 			fmt.Println(rowHead.identifier)

		// 		// curr := rowHead.right
		// 		// for curr != rowHead {
		// 		//   fmt.Print(curr.identifier)
		// 		//   fmt.Print(" -> ")
		// 		//   curr = curr.right
		// 		// }
		// 		// fmt.Println(row)
		// 		}
		// 	}
		// }
    
		fmt.Println("solutions", dlx.solutions)
		fmt.Println("expected", test.solutions)
	}

	fmt.Println("DONE")

	// solutions := dlx.Solve()
	// fmt.Println("solutions 1", solutions)

	// dlx2 := NewDlx([]string{"1", "2", "3", "4", "5", "6", "7"})
	// dlx2.AddConstraintRow([]string{"1", "4", "7"})      // a
	// dlx2.AddConstraintRow([]string{"1", "4"})           //b
	// dlx2.AddConstraintRow([]string{"4", "5", "7"})      //c
	// dlx2.AddConstraintRow([]string{"3", "5", "6"})      //d
	// dlx2.AddConstraintRow([]string{"2", "3", "6", "7"}) //e
	// dlx2.AddConstraintRow([]string{"2", "7"})           //f

	//  solutions2 := dlx2.Solve()
	// fmt.Println("solutions 2", solutions2)

}
