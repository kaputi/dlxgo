package dlx

import (
	"fmt"

	"github.com/kaputi/dlxgo/stack"
)

type Node struct {
	left, right, up, down, colHead, rowHead *Node
	identifier                              string
	colSize                                 int
	rowIndex                                int
}

func newNode() *Node {
	node := &Node{}
	node.left = node
	node.right = node
	node.up = node
	node.down = node
	node.colHead = node
	node.rowHead = node

	return node
}

type Dlx struct {
	root            *Node
	nodeCount       int
	colHeads        map[string]*Node
	rowHeads        map[int]*Node
	rowCounter      int
	removalStack    stack.Stack
	solutions       [][]int
	partialSolution []int
	solutionsFound  int
}

func NewDlx(identifiers []string) *Dlx {
	root := newNode()
	dlx := &Dlx{
		colHeads: make(map[string]*Node),
		rowHeads: make(map[int]*Node),
		root:     root,
	}

	for _, identifier := range identifiers {
		node := newNode()
		node.identifier = identifier
		dlx.colHeads[identifier] = node

		node.left = root.left
		node.right = root
		root.left.right = node
		root.left = node
	}

	return dlx
}

func (d *Dlx) AddConstraintRow(identifiers []string) {
	rowHead := newNode()
	rowHead.identifier = fmt.Sprintf("row%d", d.rowCounter)
	d.rowHeads[d.rowCounter] = rowHead

	for _, identifier := range identifiers {
		node := newNode()
		node.identifier = identifier
		node.rowIndex = d.rowCounter
		node.rowHead = rowHead
		node.left = rowHead.left
		node.right = rowHead
		rowHead.left.right = node
		rowHead.left = node

		colHead := d.colHeads[identifier]
		node.colHead = colHead
		node.down = colHead
		node.up = colHead.up
		colHead.up.down = node
		colHead.up = node

		colHead.colSize++
		d.nodeCount++
	}

	d.rowCounter++
}

func (d *Dlx) Solve2() bool {
	minColHead := d.getMinColumn()

	// fmt.Println("MIN COL HEAD", minColHead.identifier)
	// fmt.Println("MIN COL SIZE", minColHead.colSize)

	if minColHead == d.root {
		// matrix is empty = solution found
		d.solutions = append(d.solutions, d.partialSolution)
		d.restoreMatrix()
		d.partialSolution = []int{}
		return true
	}

	if minColHead.colSize == 0 && minColHead != d.root {
		d.partialSolution = []int{}
		return false
	}

	selectedRow := minColHead.down

	for selectedRow != minColHead {
		if d.nodeCount > 0 {
			d.coverRow(selectedRow)
			d.partialSolution = append(d.partialSolution, selectedRow.rowIndex)
			// fmt.Println("ADDING TO PARTIAL SOLUTION", selectedRow.rowIndex)
		}
		if !d.Solve2() {
			d.restoreMatrix()
		}
		selectedRow = selectedRow.down
	}

	d.restoreMatrix()

	return true
}

// func (d *Dlx) Solve() []solution {
// 	d.solveHelper(false)
// 	return d.solutions
// }

// func (d *Dlx) solveHelper(nested bool) {
// 	if d.nodeCount == 0 {
// 		d.solutionsFound++
// 		d.solutions = append(d.solutions, solution{})
// 		return
// 	}

// 	if d.solutionsFound == len(d.solutions) && !nested {
// 		d.solutions = append(d.solutions, solution{})
// 	}

// 	colHead := d.getMinColumn()

// 	if colHead.colSize == 0 {
// 		// no solutions
// 		d.restoreMatrix()
// 		return
// 	}

// 	selectedRow := colHead.down

// 	for selectedRow != colHead {
// 		solutionRow := solutionRow{}

// 		rowHead := selectedRow.rowHead

// 		// get colHeads in row
// 		colHeads := make([]*Node, 0)
// 		curr := rowHead.right
// 		for curr != rowHead {
// 			colHeads = append(colHeads, curr.colHead)
// 			solutionRow = append(solutionRow, curr.identifier)
// 			curr = curr.right
// 		}

// 		d.solutions[len(d.solutions)-1] = append(d.solutions[len(d.solutions)-1], solutionRow)

// 		for _, colHead := range colHeads {
// 			d.coverColumn(colHead)
// 		}

// 		if d.nodeCount != 0 {
// 			d.solveHelper(true)
// 		}

// 		d.restoreMatrix()

// 		selectedRow = selectedRow.down
// 	}
// }

func (d *Dlx) removeNode(node *Node) {
	node.left.right = node.right
	node.right.left = node.left
	node.up.down = node.down
	node.down.up = node.up

	if node.colHead != node {
		node.colHead.colSize--
		d.nodeCount--
	}

	d.removalStack.Push(node)
}

func (d *Dlx) reinsertNode(node *Node) {
	node.left.right = node
	node.right.left = node
	node.up.down = node
	node.down.up = node

	if node.colHead != node {
		node.colHead.colSize++
		d.nodeCount++
	}
}

func (d *Dlx) restoreMatrix() {
	for d.removalStack.Length > 0 {
		node := d.removalStack.Pop().(*Node)
		d.reinsertNode(node)
	}
}

func (d *Dlx) getMinColumn() *Node {
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

func (d *Dlx) coverRow(nodeInRow *Node) {
	rowHead := nodeInRow.rowHead
	curr := rowHead.right
	for curr != rowHead {
		currColHead := curr.colHead
		node := currColHead.down
		for node != currColHead {
			d.removeNode(node)
			if node.rowHead != rowHead {
				innerNode := node.rowHead.right
				for innerNode != innerNode.rowHead {
					d.removeNode(innerNode)
					innerNode = innerNode.right
				}
			}
			node = node.down
		}
		d.removeNode(currColHead)
		curr = curr.right
	}
}

// func (d *Dlx) coverColumn(colHead *Node) {
// 	curr := colHead.down
// 	for curr != colHead {
// 		currRow := curr.rowHead.right
// 		for currRow != curr.rowHead {
// 			d.removeNode(currRow)
// 			currRow = currRow.right
// 		}
// 		curr = curr.down
// 	}
// 	d.removeNode(colHead)
// }
