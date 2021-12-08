package utils

// value also indicates presence here
type Set[T comparable] map[T]bool

func (s Set[T]) Add(element T) bool {
	for v := range s {
		if element == v {
			return false
		}
	}
	s[element] = true
	return true
}

func (s Set[T]) AddAll(elements []T) bool {
	addedAtLeastOne := false
	for _, e := range elements {
		addedAtLeastOne = s.Add(e)
	}
	return addedAtLeastOne
}

func (s Set[T]) Contains(e T) bool {
	return s[e]
}

func (s Set[T]) Intersection(other Set[T]) (result Set[T]) {
	result = make(Set[T], 0)
	for v := range s {
		if other[v] {
			result[v] = true
		}
	}
	return
}

func (s Set[T]) Union(other Set[T]) (result Set[T]) {
	result = make(Set[T], 0)

	for v := range s {
		result[v] = true
	}

	for v := range other {
		result[v] = true
	}

	return
}

func (s Set[T]) Difference(other Set[T]) (result Set[T]) {
	result = make(Set[T], 0)

	for v := range s {
		if !other[v] {
			result[v] = true
		}
	}
	return
}

func (s Set[T]) Equals(other Set[T]) bool {
	for v := range s {
		if !other[v] {
			return false
		}
	}

	for v := range other {
		if !s[v] {
			return false
		}
	}
	return true
}

func (s Set[T]) Values() (values []T) {
	values = make([]T, 0)
	for v := range s {
		values = append(values, v)
	}
	return
}

func SetFrom[T comparable](s []T) (result Set[T]) {
	result = make(Set[T], 0)

	for _, val := range s {
		result[val] = true
	}
	return
}
