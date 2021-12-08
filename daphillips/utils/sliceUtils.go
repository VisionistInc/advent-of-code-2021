package utils

import "constraints"

// wanted to try using a type, but couldn't get it to go today
// type Bunch[T any] []T

func Map[T any, F any](src []T, mapper func(in T) F) (result []F) {
	for _, item := range src {
		result = append(result, mapper(item))
	}
	return
}

func Filter[T any](src []T, isValid func(in T) bool) (result []T) {
	for _, item := range src {
		if isValid(item) {
			result = append(result, item)
		}
	}
	return
}

func Reduce[T constraints.Integer](src []T, initialVal T, reducer func(acc T, val T) T) (result T) {
	result = initialVal
	for _, next := range src {
		result = reducer(result, next)
	}
	return
}
