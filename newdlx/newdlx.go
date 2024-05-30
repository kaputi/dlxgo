package newdlx

import (
	"fmt"
)

type node struct {
	up, down, left, right, column *node
	columnSize                    int
	identifier                    string
}

type dlxMatrix struct {
	root            *node
	partialSolution [][]string
	solutions       [][][]string
	columns         map[string]*node
	debug           bool
}

func NewDlxMatrix(identifiers []string) *dlxMatrix {
	root := &node{}
	root.up = root
	root.down = root
	root.left = root
	root.right = root

	dlx := &dlxMatrix{
		root:    root,
		columns: make(map[string]*node),
	}

	for _, identifier := range identifiers {
		node := &node{identifier: identifier}
		dlx.columns[identifier] = node
		node.column = node
		node.up = node
		node.down = node
		node.left = root.left
		node.right = root
		root.left.right = node
		root.left = node
	}

	return dlx
}

func (d *dlxMatrix) AddConstraintRow(row []string) {
	firstNode := &node{identifier: row[0]}
	addNodeToCol(d.columns[row[0]], firstNode)
	firstNode.left = firstNode
	firstNode.right = firstNode

	for _, identifier := range row[1:] {
		node := &node{identifier: identifier}
		addNodeToCol(d.columns[identifier], node)
		node.right = firstNode
		node.left = firstNode.left
		firstNode.left.right = node
		firstNode.left = node
	}
}

func addNodeToCol(column, node *node) {
	node.column = column
	node.up = column.up
	node.down = column
	column.up.down = node
	column.up = node
	column.columnSize++
}

func (d *dlxMatrix) getSmallestCol() *node {
	curr := d.root.right
	min := curr
	for curr != d.root {
		if curr.columnSize < min.columnSize {
			min = curr
		}
		curr = curr.right
	}
	return min
}

func (d *dlxMatrix) removeNode(node *node) {
	// fmt.Printf("node %v right now points to %v\n", node.left.identifier, node.right.identifier)
	node.left.right = node.right
	// fmt.Printf("node %v left now points to %v\n", node.right.identifier, node.left.identifier)
	node.right.left = node.left
	// fmt.Printf("node %v down now points to %v\n", node.up.identifier, node.down.identifier)
	node.up.down = node.down
	// fmt.Printf("node %v up now points to %v\n", node.down.identifier, node.up.identifier)
	node.down.up = node.up
	if node != node.column {
		node.column.columnSize--
	}
}

func (d *dlxMatrix) restoreNode(node *node) {
	node.left.right = node
	node.right.left = node
	node.up.down = node
	node.down.up = node
	if node != node.column {
		node.column.columnSize++
	}
}

func (d *dlxMatrix) getRowIdentifiers(node *node) []string {
	identifiers := []string{node.identifier}
	curr := node.right
	for curr != node {
		identifiers = append(identifiers, curr.identifier)
		curr = curr.right
	}
	return identifiers
}

func (d *dlxMatrix) coverColumn(columns []*node) {
	for _, column := range columns {
		// remove from header the header
		column.left.right = column.right
		column.right.left = column.left

		curr := column.down
		for curr != column {
			node := curr.right
			for node != curr {
				d.removeNode(node)
				node = node.right
			}
			curr = curr.down
		}
	}
}

func (d *dlxMatrix) uncoverColumn(columns []*node) {
	for _, column := range columns {
		curr := column.up
		for curr != column {
			node := curr.left
			for node != curr {
				d.restoreNode(node)
				node = node.left
			}
			curr = curr.up
		}

		column.left.right = column
		column.right.left = column
	}
}

func (d *dlxMatrix) SolveOne() {
	count := 0
	d.solve(false, &count)
	// TODO:  return solution somehow
}

func (d *dlxMatrix) SolveAll() {
	level := 0
	d.solve(true, &level)
	// TODO:  return solution somehow
}

func (d *dlxMatrix) printColumnsInMatrix() {
	testN := d.root.right
	fmt.Print("Columns in matrix: ")
	for testN != d.root {
		fmt.Printf("%v size: %d | ", testN.identifier, testN.columnSize)
		testN = testN.right
	}
	fmt.Println()
}

func (d *dlxMatrix) solve(multiple bool, level *int) bool {
	fmt.Printf("Solving Level: %d ----------", *level)
	fmt.Println()

	if d.root.right == d.root {
		// matrix is empty so solution is found
		fmt.Println("solution found at level: ", *level)
		fmt.Println("current partSolution: ", d.partialSolution)
		// if len(d.partialSolution) > 0 {
		d.solutions = append(d.solutions, d.partialSolution)
		// }
		fmt.Println("new partSolution: ", d.partialSolution)
		fmt.Println()
		return true
	}

	smallestCol := d.getSmallestCol()

	if smallestCol.columnSize == 0 {
		fmt.Printf("no solution found, column %v has size: %d\n", smallestCol.identifier, smallestCol.columnSize)
		fmt.Println("Returning false from level: ", *level)
		fmt.Println()
		// no solution found
		fmt.Println("reseting partial solution")
		fmt.Println("partialSolution: ", d.partialSolution)
		fmt.Println()
		d.partialSolution = [][]string{}
		return false
	}

	selectedRowNode := smallestCol.down
	for selectedRowNode != smallestCol {

		d.printColumnsInMatrix()

		selectedRowIdentifiers := d.getRowIdentifiers(selectedRowNode)
		fmt.Println("columnsToCover: ", selectedRowIdentifiers)
		fmt.Println("adding row to partial solution: ", selectedRowIdentifiers)
		fmt.Println()
		// TODO: to get only one solution, loop until solution is found, then break loop

		d.partialSolution = append(d.partialSolution, d.getRowIdentifiers(selectedRowNode))

		columnsToCover := []*node{selectedRowNode.column}

		node := selectedRowNode.right
		for node != selectedRowNode {
			columnsToCover = append(columnsToCover, node.column)
			node = node.right
		}

		d.coverColumn(columnsToCover)

		*level++
		solutionFound := d.solve(multiple, level)
		*level--

		if solutionFound && !multiple {
			fmt.Println("Solution found, and multiple is false")
			fmt.Println("Returning true from level: ", *level)
			fmt.Println()
			return true
		}

		fmt.Println("uncovering columns: ", selectedRowIdentifiers)
		d.uncoverColumn(columnsToCover)
		d.printColumnsInMatrix()
		fmt.Print("removing row from partial solution: ", selectedRowIdentifiers)
		d.partialSolution = d.partialSolution[:len(d.partialSolution)-1]
		fmt.Println()

		fmt.Println("moving to next row at level: ", *level)
		selectedRowNode = selectedRowNode.down
	}

	return true
}
