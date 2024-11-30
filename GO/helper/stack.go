package helper

type Stack[T any] []T

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	size := len(*s)
	item := (*s)[size-1]
	*s = (*s)[:size-1]
	return item
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
