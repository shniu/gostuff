package stack

// ArrayStack 的通用实现
// 顺序栈的实现
type ArrayStack struct {
	items    []string
	capacity int // capacity of stack
	size     int // size of stack elements
}

func NewStack(capacity int) *ArrayStack {
	return &ArrayStack{
		capacity: capacity,
		items:    make([]string, capacity),
	}
}

func (s *ArrayStack) isFull() bool {
	return s.size >= s.capacity
}

func (s *ArrayStack) isEmpty() bool {
	return s.size <= 0
}

// 入栈
func (s *ArrayStack) Push(item string) bool {
	// 满了直接返回false
	if s.isFull() { return false }

	s.items[s.size] = item
	s.size++
	return true
}

// 出栈
func (s *ArrayStack) Pop() string {
	if s.isEmpty() { return "" }

	top := s.items[s.size - 1]
	s.size--
	return top
}

// peek
func (s *ArrayStack) Peek() string {
	if s.isEmpty() { return "" }

	return s.items[s.size - 1]
}

// Len
func (s *ArrayStack) Len() int {
	return s.size
}
