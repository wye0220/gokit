package set

// Empty is an alias for the empty struct used in the map.
type Empty struct{}

// Set is a generic interface for a set data structure.
type Set[T comparable] interface {
	Add(key T)
	Remove(key T)
	Contains(key T) bool
	Len() int
	Clear()
	Keys() []T
}

// HashSet is a hash-based implementation of the Set interface.
type HashSet[T comparable] struct {
	m map[T]Empty
}

// New creates and returns a new HashSet with the specified initial size.
func New[T comparable](size int) *HashSet[T] {
	return &HashSet[T]{
		m: make(map[T]Empty, size),
	}
}

// Add inserts the given key into the set.
func (s *HashSet[T]) Add(key T) {
	s.m[key] = Empty{}
}

// Remove deletes the given key from the set.
func (s *HashSet[T]) Remove(key T) {
	delete(s.m, key)
}

// Contains checks if the given key is in the set.
func (s *HashSet[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Len returns the number of elements in the set.
func (s *HashSet[T]) Len() int {
	return len(s.m)
}

// Clear removes all elements from the set.
func (s *HashSet[T]) Clear() {
	s.m = make(map[T]Empty)
}

// Keys returns a slice of all keys in the set.
func (s *HashSet[T]) Keys() []T {
	return mapKeys(s.m)
}

// mapKeys extracts the keys from the map and returns them as a slice.
func mapKeys[T comparable](m map[T]Empty) []T {
	res := make([]T, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}
