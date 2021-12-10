package utils

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(next T) {
	*s = append(*s, next)
}

// TODO I don't want to have to panic here... but in order to get that I'd need to be able to assign nil on return, which i can only do with a pointer
func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("popping on empty stack!")
	}

	lastIndex := len(*s) - 1
	lastElement := (*s)[lastIndex]
	*s = (*s)[:lastIndex]

	return lastElement

}

// TODO I don't want to have to panic here... but in order to get that I'd need to be able to assign nil on return, which i can only do with a pointer
func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("peeking on empty stack!")
	}
	return (*s)[len(*s)-1]
}
