package calendar

type Stack struct {
	queue []*Board
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(bs ...*Board) {
	s.queue = append(s.queue, bs...)
}

func (s *Stack) Pop() *Board {
	end := len(s.queue) - 1
	result := s.queue[end]
	s.queue = s.queue[:end]
	return result
}

func (s *Stack) IsEmpty() bool {
	return len(s.queue) == 0
}
