package ddlx

type StackNode struct {
	prev *StackNode
	item *DNode
}

type Stack struct {
	head   *StackNode
	Length int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(item *DNode) {
	s.Length++
	node := &StackNode{item: item}

	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack) pop() *DNode {
	s.Length = max(0, s.Length-1)

	if s.Length == 0 {
		item := s.head.item
		s.head = nil
		return item
	}

	head := s.head
	s.head = head.prev

	return head.item
}

func (s *Stack) peek() *DNode {
	if s.head == nil {
		return nil
	}
	return s.head.item
}

func (s *Stack) toSlice() []*DNode {
	slice := make([]*DNode, s.Length)
	curr := s.head

	for i := s.Length - 1; i >= 0; i-- {
		slice[i] = curr.item
		curr = curr.prev
	}

	return slice
}

func (s *Stack) empty() {
	s.Length = 0
	s.head = nil
}
