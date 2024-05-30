package newdlx

import (
	"fmt"
)

type dlxNode struct {
	up, down, left, right, column *dlxNode
	columnSize                    int
	identifier                    string
}

type dlxMatrix struct {
	root            *dlxNode
	partialSolution [][]string
	solutions       [][][]string
	columns         map[string]*dlxNode
	debug           bool
}

func NewDlxMatrix(identifiers []string) *dlxMatrix {
	root := &dlxNode{}
	root.up = root
	root.down = root
	root.left = root
	root.right = root

	dlx := &dlxMatrix{
		root:    root,
		columns: make(map[string]*dlxNode),
	}

	for _, identifier := range identifiers {
		node := &dlxNode{identifier: identifier}
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
	firstNode := &dlxNode{identifier: row[0]}
	addNodeToCol(d.columns[row[0]], firstNode)
	firstNode.left = firstNode
	firstNode.right = firstNode

	for _, identifier := range row[1:] {
		node := &dlxNode{identifier: identifier}
		addNodeToCol(d.columns[identifier], node)
		node.right = firstNode
		node.left = firstNode.left
		firstNode.left.right = node
		firstNode.left = node
	}
}

func addNodeToCol(column, node *dlxNode) {
	node.column = column
	node.up = column.up
	node.down = column
	column.up.down = node
	column.up = node
	column.columnSize++
}

func (d *dlxMatrix) getSmallestCol() *dlxNode {
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

func (d *dlxMatrix) removeNode(node *dlxNode) {
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

func (d *dlxMatrix) restoreNode(node *dlxNode) {
	node.left.right = node
	node.right.left = node
	node.up.down = node
	node.down.up = node
	if node != node.column {
		node.column.columnSize++
	}
}

func (d *dlxMatrix) getRowIdentifiers(node *dlxNode) []string {
	identifiers := []string{node.identifier}
	curr := node.right
	for curr != node {
		identifiers = append(identifiers, curr.identifier)
		curr = curr.right
	}
	return identifiers
}

func (d *dlxMatrix) coverColumn(columns []*dlxNode) {
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

func (d *dlxMatrix) uncoverColumn(columns []*dlxNode) {
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
	level := 0
	d.solve(false, &level)
	// TODO:  return solution somehow
}

func (d *dlxMatrix) SolveAll() {
	level := 0
	d.solve(true, &level)
	// TODO:  return solution somehow
}

func (d *dlxMatrix) getColsInMatrix() string {
	str := "Columns in matrix: "
	testN := d.root.right
	for testN != d.root {
		str += fmt.Sprintf("%v size: %d | ", testN.identifier, testN.columnSize)
		testN = testN.right
	}
	return str
}

func (d *dlxMatrix) solve(multiple bool, level *int) bool {
	d.logAtLevel(*level, fmt.Sprintf("Solving Level %d ----------\n", *level))

	if d.root.right == d.root {
		// matrix is empty so solution is found
		d.logAtLevel(*level, fmt.Sprintf("*** Solution found at level %d ***\n", *level))
		d.solutions = append(d.solutions, d.partialSolution)
		d.logAtLevel(*level, fmt.Sprintf("Adding partial solution %v to solutions\n", d.partialSolution))
		d.logAtLevel(*level, fmt.Sprintf("current solutions: %v\n", d.solutions))
		// d.partialSolution = [][]string{}
		// fmt.Println("reseting part solutions: ", d.partialSolution)
		// fmt.Println("new partSolution: ", d.partialSolution)
		d.logln("")
		return true
	}

	smallestCol := d.getSmallestCol()

	if smallestCol.columnSize == 0 {
		d.logAtLevel(*level, fmt.Sprintf("no solution found, column %v has size: %d\n", smallestCol.identifier, smallestCol.columnSize))
		d.logAtLevel(*level, fmt.Sprintf("Returning false from level: %d ", *level))
		d.logln("")
		// no solution found
		d.logAtLevel(*level, "reseting partial solution")
		d.logAtLevel(*level, fmt.Sprintf("partialSolution: %v", d.partialSolution))
		d.logln("")
		d.partialSolution = [][]string{}
		return false
	}

	selectedRowNode := smallestCol.down
	for selectedRowNode != smallestCol {

		d.logAtLevel(*level, d.getColsInMatrix())
		d.logln("")

		selectedRowIdentifiers := d.getRowIdentifiers(selectedRowNode)
		d.logAtLevel(*level, fmt.Sprintf("columnsToCover: %v\n", selectedRowIdentifiers))
		d.logAtLevel(*level, fmt.Sprintf("adding row to partial solution: %v\n\n", selectedRowIdentifiers))
		// TODO: to get only one solution, loop until solution is found, then break loop

		d.partialSolution = append(d.partialSolution, d.getRowIdentifiers(selectedRowNode))

		columnsToCover := []*dlxNode{selectedRowNode.column}
		d.logAtLevel(*level, fmt.Sprintf("adding %v to columnsToCover\n", selectedRowNode.column.identifier))

		node := selectedRowNode.right
		// fmt.Printf("AQUI %v, %v\n", selectedRowNode.identifier, node.identifier)
		for node != selectedRowNode {
			columnsToCover = append(columnsToCover, node.column)
			d.logAtLevel(*level, fmt.Sprintf("adding %v to columnsToCover, in loop\n", node.identifier))
			node = node.right
		}

		d.logAtLevel(*level, "covering columns: [")
		for _, column := range columnsToCover {
			d.log(fmt.Sprintf(" %v", column.identifier))
		}
		d.log(" ]\n\n")

		d.coverColumn(columnsToCover)

		*level++
		solutionFound := d.solve(multiple, level)
		*level--

		if solutionFound && !multiple {
			d.logAtLevel(*level, "Solution found, and multiple is false\n")
			d.logAtLevel(*level, fmt.Sprintf("Returning true from level: %d\n\n", *level))
			return true
		}

		d.logAtLevel(*level, fmt.Sprintf("uncovering columns: %v\n", selectedRowIdentifiers))
		d.uncoverColumn(columnsToCover)
		d.getColsInMatrix()
		d.logAtLevel(*level, fmt.Sprintf("removing row from partial solution: %v\n", selectedRowIdentifiers))
		if len(d.partialSolution) > 0 {
			d.partialSolution = d.partialSolution[:len(d.partialSolution)-1]
		}
		d.logAtLevel(*level, fmt.Sprintf("partialSolution: %v\n", d.partialSolution))

		if selectedRowNode.down != smallestCol {
			d.logAtLevel(*level, fmt.Sprintf("moving to next row at level: %d\n", *level))
		}

		selectedRowNode = selectedRowNode.down
	}

	return true
}

func (d *dlxMatrix) log(msg string) {
	if d.debug {
		fmt.Print(msg)
	}
}

func (d *dlxMatrix) logln(msg string) {
	if d.debug {
		fmt.Println(msg)
	}
}

func (d *dlxMatrix) logAtLevel(level int, msg string) {
	tab := " ••••"
	fullTab := ""
	for i := 0; i < level; i++ {
		fullTab += tab
	}
	fmt.Printf("%s %s", fullTab, msg)
}
