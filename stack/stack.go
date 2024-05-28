package stack

type StackNode struct {
	prev *StackNode
	item interface{}
}

type Stack struct {
	head   *StackNode
	Length int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item interface{}) {
	s.Length++
	node := &StackNode{item: item}

	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack) Pop() interface{} {
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

func (s *Stack) Peek() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.item
}
