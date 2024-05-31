package ddlx

import (
	"fmt"
)

type DNode struct {
	up, down, left, right *DNode
	colHead               *DNode
	rowHead               *DNode
	identifier            string
	colSize               int
}

type Dlx struct {
	root            *DNode
	partialSolution *Stack
	solutions       [][]*DNode
	columns         map[string]*DNode
}

func NewDlx(colNames []string) *Dlx {
	root := &DNode{identifier: "root"}
	root.up = root
	root.down = root
	root.left = root
	root.right = root

	dlx := &Dlx{
		root:            root,
		columns:         make(map[string]*DNode),
		partialSolution: newStack(),
	}

	for _, name := range colNames {
		node := &DNode{
			identifier: name,
		}

		dlx.columns[name] = node

		node.colHead = node

		node.up = node
		node.down = node
		node.left = dlx.root.left
		node.right = dlx.root

		dlx.root.left.right = node
		dlx.root.left = node
	}

	return dlx
}

func (d *Dlx) AddConstraintRow(rowName string, colNames []string) {
	rowHead := &DNode{identifier: rowName}

	rowHead.left = rowHead
	rowHead.right = rowHead

	for _, colName := range colNames {
		colHead := d.columns[colName]

		colHead.colSize++

		node := &DNode{
			identifier: fmt.Sprintf("%s_%s", rowName, colName),
			left:       rowHead.left,
			right:      rowHead,
			up:         colHead.up,
			down:       colHead,
			rowHead:    rowHead,
			colHead:    colHead,
		}

		rowHead.left.right = node
		rowHead.left = node

		colHead.up.down = node
		colHead.up = node

	}
}

func (d *Dlx) SolveOne() {
	d.solve(false)
}

func (d *Dlx) SolveAll() [][]string {
	d.solve(true)

	solutions := make([][]string, len(d.solutions))

	for i, solution := range d.solutions {
		solutions[i] = make([]string, len(solution))
		for j, rowHead := range solution {
			solutions[i][j] = rowHead.identifier
		}
	}

	return solutions
}

func (d *Dlx) GetColNamesFromSolutions() [][][]string {
	solutions := make([][][]string, len(d.solutions))

	for i, solution := range d.solutions {
		solutions[i] = make([][]string, len(solution))
		for j, rowHead := range solution {
			colNames := []string{}
			curr := rowHead.right
			for curr != rowHead {
				colNames = append(colNames, curr.colHead.identifier)
				curr = curr.right
			}
			solutions[i][j] = colNames
		}
	}

	return solutions
}

func (d *Dlx) getSmallestCol() *DNode {
	curr := d.root.right
	min := curr

	for curr != d.root {
		if curr.colSize < min.colSize {
			min = curr
		}
		curr = curr.right
	}

	return min
}

func (d *Dlx) removeNode(node *DNode) {
	node.left.right = node.right
	node.right.left = node.left
	node.up.down = node.down
	node.down.up = node.up

	if node != node.colHead {
		node.colHead.colSize--
	}
}

func (d *Dlx) restoreNode(node *DNode) {
	node.left.right = node
	node.right.left = node
	node.up.down = node
	node.down.up = node

	if node != node.colHead {
		node.colHead.colSize++
	}
}

func (d *Dlx) cover(columns []*DNode) *Stack {
	coveredNodes := newStack()

	for _, col := range columns {
		// make column inaccesible from head
		col.left.right = col.right
		col.right.left = col.left

		coveredNodes.push(col)

		curr := col.down

		for curr != col {
			node := curr.rowHead.right
			for node != curr.rowHead {
				d.removeNode(node)
				coveredNodes.push(node)
				node = node.right
			}
			curr = curr.down
		}
	}

	return coveredNodes
}

func (d *Dlx) uncover(covered *Stack) {
	for covered.Length > 0 {
		node := covered.pop()
		d.restoreNode(node)
	}
}

func (d *Dlx) solve(multiple bool) bool {
	if d.root.right == d.root {
		d.solutions = append(d.solutions, d.partialSolution.toSlice())
		return true
	}

	smallestCol := d.getSmallestCol()
	if smallestCol.colSize == 0 {
		d.partialSolution.empty()
		return false
	}

	selectedRowNode := smallestCol.down
	for selectedRowNode != smallestCol {
		d.partialSolution.push(selectedRowNode.rowHead)

		columnsToCover := []*DNode{}
		// we start always from the head
		curr := selectedRowNode.rowHead.right
		for curr != selectedRowNode.rowHead {
			columnsToCover = append(columnsToCover, curr.colHead)
			curr = curr.right
		}

		coveredColumns := d.cover(columnsToCover)

		solutionFound := d.solve(multiple)

		d.uncover(coveredColumns)

		if solutionFound && !multiple {
			return true
		}

		if d.partialSolution.Length > 0 {
			d.partialSolution.pop()
		}

		selectedRowNode = selectedRowNode.down
	}

	return true
}
