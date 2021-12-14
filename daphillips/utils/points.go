package utils

import "constraints"

// todo template it ?
type Point[T constraints.Ordered] struct {
	X T
	Y T
}

func (p Point[T]) Translate(dx, dy T) Point[T] {
	return Point[T]{p.X + dx, p.Y + dy}
}
