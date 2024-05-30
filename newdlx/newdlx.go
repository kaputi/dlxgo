package newdlx

import (
	"fmt"

	"github.com/kaputi/dlxgo/stack"
)

type dlxNode struct {
	up, down, left, right, column *dlxNode
	rowIdentifier                 string
	rowNumber                     int
	columnSize                    int
	colIdentifier                 string
}

type dlxMatrix struct {
	root            *dlxNode
	partialSolution [][]string
	solutions       [][][]string
	columns         map[string]*dlxNode
	rowCounter      int
	debug           bool
}

func NewDlxMatrix(identifiers []string) *dlxMatrix {
	root := &dlxNode{colIdentifier: "root", rowIdentifier: "root"}
	root.up = root
	root.down = root
	root.left = root
	root.right = root

	dlx := &dlxMatrix{
		root:    root,
		columns: make(map[string]*dlxNode),
	}

	for _, identifier := range identifiers {
		node := &dlxNode{
			colIdentifier: identifier,
			rowIdentifier: "head",
		}
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

func (d *dlxMatrix) AddConstraintRow(rowIdentifier string, row []string) {
	firstNode := &dlxNode{
		colIdentifier: row[0],
		rowIdentifier: rowIdentifier,
		rowNumber:     d.rowCounter,
	}
	addNodeToCol(d.columns[row[0]], firstNode)
	firstNode.left = firstNode
	firstNode.right = firstNode

	for _, identifier := range row[1:] {
		node := &dlxNode{
			colIdentifier: identifier,
			rowIdentifier: rowIdentifier,
			rowNumber:     d.rowCounter,
		}
		addNodeToCol(d.columns[identifier], node)
		node.right = firstNode
		node.left = firstNode.left
		firstNode.left.right = node
		firstNode.left = node
	}
	d.rowCounter++
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
	d.logln(fmt.Sprintf("Removing Node %v,%v", node.rowIdentifier, node.colIdentifier))
	d.logln(fmt.Sprintf(" Node %v,%v right -> %v,%v",
		node.left.rowIdentifier,
		node.left.colIdentifier,
		node.right.rowIdentifier,
		node.right.colIdentifier,
	))
	node.left.right = node.right

	d.logln(fmt.Sprintf(" Node %v,%v left -> %v,%v",
		node.right.rowIdentifier,
		node.right.colIdentifier,
		node.left.rowIdentifier,
		node.left.colIdentifier,
	))
	node.right.left = node.left

	d.logln(fmt.Sprintf(" Node %v,%v down -> %v,%v",
		node.up.rowIdentifier,
		node.up.colIdentifier,
		node.down.rowIdentifier,
		node.down.colIdentifier,
	))
	node.up.down = node.down

	d.logln(fmt.Sprintf(" Node %v,%v up -> %v,%v",
		node.down.rowIdentifier,
		node.down.colIdentifier,
		node.up.rowIdentifier,
		node.up.colIdentifier,
	))
	node.down.up = node.up

	if node != node.column {
		node.column.columnSize--
	}

}

func (d *dlxMatrix) restoreNode(node *dlxNode) {
	d.logln(fmt.Sprintf("RESTORING NODE %v,%v =====================", node.rowIdentifier, node.colIdentifier))

	d.logln(fmt.Sprintf(" Node %v,%v right -> %v,%v",
		node.left.rowIdentifier,
		node.left.colIdentifier,
		node.rowIdentifier,
		node.colIdentifier,
	))
	node.left.right = node

	d.logln(fmt.Sprintf(" Node %v,%v left -> %v,%v",
		node.right.rowIdentifier,
		node.right.colIdentifier,
		node.rowIdentifier,
		node.colIdentifier,
	))
	node.right.left = node

	d.logln(fmt.Sprintf(" Node %v,%v down -> %v,%v",
		node.up.rowIdentifier,
		node.up.colIdentifier,
		node.rowIdentifier,
		node.colIdentifier,
	))
	node.up.down = node

	d.logln(fmt.Sprintf(" Node %v,%v up -> %v,%v",
		node.down.rowIdentifier,
		node.down.colIdentifier,
		node.rowIdentifier,
		node.colIdentifier,
	))
	node.down.up = node
	if node != node.column {
		node.column.columnSize++
	}
}

func (d *dlxMatrix) getRowIdentifiers(node *dlxNode) []string {
	identifiers := []string{node.colIdentifier}
	curr := node.right
	for curr != node {
		identifiers = append(identifiers, curr.colIdentifier)
		curr = curr.right
	}
	return identifiers
}

func (d *dlxMatrix) coverColumn(columns []*dlxNode) stack.Stack {
	coveredNodes := stack.NewStack()
	for _, column := range columns {
		// remove from header the header
		d.logln(fmt.Sprintf("COVERING COLUMN %v =====================", column.colIdentifier))
		d.logln("Removing Head")
		d.logln(fmt.Sprintf(" %v right -> %v",
			column.left.colIdentifier,
			column.right.colIdentifier,
		))
		d.logln(fmt.Sprintf(" %v left -> %v",
			column.right.colIdentifier,
			column.left.colIdentifier,
		))
		column.left.right = column.right
		column.right.left = column.left
		coveredNodes.Push(column)

		curr := column.down
		for curr != column {
			node := curr.right
			for node != curr {
				d.removeNode(node)
				coveredNodes.Push(node)
				node = node.right
			}
			curr = curr.down
		}

		d.logln("")
	}

	return *coveredNodes
}

func (d *dlxMatrix) uncover(coveredNodes stack.Stack) {
	for coveredNodes.Length > 0 {
		node := coveredNodes.Pop().(*dlxNode)
		d.restoreNode(node)
	}
}

func (d *dlxMatrix) uncoverColumn(columns []*dlxNode) {
	// for _, column := range columns {
	for i := len(columns) - 1; i >= 0; i-- {
		column := columns[i]
		d.logln(fmt.Sprintf("RESTORING COLUMN %v =====================", column.colIdentifier))
		d.logln(fmt.Sprintf("Col up %v,%v", column.up.rowIdentifier, column.up.colIdentifier))
		curr := column.up
		for curr != column {
			node := curr.left
			for node != curr {
				d.restoreNode(node)
				node = node.left
			}
			curr = curr.up
		}

		d.logln(" Restoring Head")
		d.logln(fmt.Sprintf(" %v right -> %v",
			column.left.colIdentifier,
			column.colIdentifier,
		))
		d.logln(fmt.Sprintf(" %v left -> %v",
			column.right.colIdentifier,
			column.colIdentifier,
		))
		column.left.right = column
		column.right.left = column
	}
	d.logln("")
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
		str += fmt.Sprintf("%v size: %d | ", testN.colIdentifier, testN.columnSize)
		testN = testN.right
	}
	return str
}

func (d *dlxMatrix) solve(multiple bool, level *int) bool {
	d.logAtLevel(*level, fmt.Sprintf("Solving Level %d ----------\n", *level))

	if d.root.right == d.root {
		// matrix is empty so solution is found
		d.logAtLevel(*level, fmt.Sprintf("*** Solution found at level %d ***\n", *level))
		currentSolution := make([][]string, len(d.partialSolution))
		copy(currentSolution, d.partialSolution)
		d.solutions = append(d.solutions, currentSolution)

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
		d.logAtLevel(*level, fmt.Sprintf("no solution found, column %v has size: %d\n", smallestCol.colIdentifier, smallestCol.columnSize))
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

		d.partialSolution = append(d.partialSolution, d.getRowIdentifiers(selectedRowNode))
		d.logAtLevel(*level, fmt.Sprintf("partialSolution: %v\n", d.partialSolution))

		columnsToCover := []*dlxNode{selectedRowNode.column}
		d.logAtLevel(*level, fmt.Sprintf("adding %v to columnsToCover\n", selectedRowNode.column.colIdentifier))

		node := selectedRowNode.right
		// fmt.Printf("AQUI %v, %v\n", selectedRowNode.identifier, node.identifier)
		for node != selectedRowNode {
			columnsToCover = append(columnsToCover, node.column)
			d.logAtLevel(*level, fmt.Sprintf("adding %v to columnsToCover, in loop\n", node.colIdentifier))
			node = node.right
		}

		d.logAtLevel(*level, "covering columns: [")
		for _, column := range columnsToCover {
			d.log(fmt.Sprintf(" %v", column.colIdentifier))
		}
		d.log(" ]\n\n")

		coveredColumns := d.coverColumn(columnsToCover)

		*level++
		solutionFound := d.solve(multiple, level)
		*level--

		if solutionFound && !multiple {
			d.logAtLevel(*level, "Solution found, and multiple is false\n")
			d.logAtLevel(*level, fmt.Sprintf("Returning true from level: %d\n\n", *level))
			return true
		}

		d.logAtLevel(*level, fmt.Sprintf("uncovering columns: %v\n", selectedRowIdentifiers))
		// d.uncoverColumn(columnsToCover)
		d.uncover(coveredColumns)
		// d.getColsInMatrix()
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
	if d.debug {
		// tab := " ••••"
		tab := " •"
		fullTab := ""
		for i := 0; i < level; i++ {
			fullTab += tab
		}
		fmt.Printf("%s %s", fullTab, msg)

	}
}
