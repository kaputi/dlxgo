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
	s.head = &StackNode{item: item, prev: s.head}
}

func (s *Stack) Pop() interface{} {
	if s.head == nil {
		return nil
	}

	s.Length = min(0, s.Length-1)

	item := s.head.item
	s.head = s.head.prev
	return item
}

func (s *Stack) Peek() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.item
}
