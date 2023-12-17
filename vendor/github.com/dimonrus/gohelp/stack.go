package gohelp

// Stack of any type
type Stack[T any] []T

// Pop value
func (s Stack[T]) Pop() (T, bool, Stack[T]) {
	if len(s) > 0 {
		return s[len(s)-1], true, s[:len(s)-1]
	}
	return *new(T), false, s
}

// Push value
func (s Stack[T]) Push(v T) Stack[T] {
	return append(s, v)
}
