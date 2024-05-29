package dlx

import (
	"fmt"
	"sort"

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
}

func NewDlx(identifiers []string) *Dlx {
	root := newNode()
	dlx := &Dlx{
		colHeads: make(map[string]*Node),
		rowHeads: make(map[int]*Node),
		root:     root,
	}

	for _, identifier := range identifiers {
		columnHead := newNode()
		columnHead.identifier = identifier
		dlx.colHeads[identifier] = columnHead

		columnHead.right = root
		columnHead.left = root.left
		root.left.right = columnHead
		root.left = columnHead
	}

	return dlx
}

func (d *Dlx) AddConstraintRow(identifiers []string) {
	rowHead := newNode()
	rowHead.identifier = fmt.Sprintf("row_%d", d.rowCounter)
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

func (d *Dlx) Solve() bool {
	minColHead := d.getMinColumn()

	if minColHead == d.root {
		// matrix is empty = solution found
		// fmt.Println("SOLUTION FOUND", d.partialSolution)
		sort.Ints(d.partialSolution)
		exists := false
		for _, solution := range d.solutions {
			if EqualSlice(solution, d.partialSolution) {
				exists = true
				break
			}
		}
		if !exists {
			d.solutions = append(d.solutions, d.partialSolution)
		}
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
			d.partialSolution = append(d.partialSolution, selectedRow.rowIndex)
			d.coverRow(selectedRow)
			// fmt.Println("ADDING TO PARTIAL SOLUTION", selectedRow.rowIndex)
		}
		if !d.Solve() {
			d.restoreMatrix()
		}
		selectedRow = selectedRow.down
	}

	d.restoreMatrix()

	return true
}

func (d *Dlx) removeNode(node *Node) {
	if node.colHead != node {
		node.colHead.colSize--
		d.nodeCount--
	}

	node.left.right = node.right
	node.right.left = node.left
	node.up.down = node.down
	node.down.up = node.up

	d.removalStack.Push(node)
}

func (d *Dlx) reinsertNode(node *Node) {
	if node.colHead != node {
		node.colHead.colSize++
		d.nodeCount++
	}

	node.left.right = node
	node.right.left = node
	node.up.down = node
	node.down.up = node
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

func (d *Dlx) PrintSolutions() {
	for _, solution := range d.solutions {
		fmt.Print("[")
		for _, row := range solution {
			fmt.Print("[")
			curr := d.rowHeads[row].right
			for curr != d.rowHeads[row] {
				fmt.Printf("%s", curr.identifier)
				if curr.right != d.rowHeads[row] {
					fmt.Print(", ")
				}
				curr = curr.right
			}
			fmt.Print("]")
			if row != solution[len(solution)-1] {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
}

func EqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
