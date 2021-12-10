package utils

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(next T) {
	*s = append(*s, next)
}

// removes the top element of the stack and returns the value in addition to a bool indicating the element was indeed in the stack
// if the stack is empty, returns the default "zero" value of T and false
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		// since there was nothing in the stack, we'll just return the default zero value of type T
		var nothing T
		return nothing, false
	}

	lastIndex := len(*s) - 1
	lastElement := (*s)[lastIndex]
	*s = (*s)[:lastIndex]

	return lastElement, true

}

// returns (but does not remove) the top element of the stack and returns the value in addition to a bool indicating the element was indeed in the stack
// if the stack is empty, returns the default "zero" value of T and false
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		// since there was nothing in the stack, we'll just return the default zero value of type T
		var nothing T
		return nothing, false
	}
	return (*s)[len(*s)-1], true
}
