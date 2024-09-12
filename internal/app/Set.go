package app

type Set[T comparable] struct {
	items map[T]bool
}

// NewSet creates a new set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

// Add inserts an item into the set
func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

// Remove deletes an item from the set
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains checks if the item is in the set
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Size returns the number of items in the set
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Values returns all the items in the set as a slice
func (s *Set[T]) Values() []T {
	values := make([]T, 0, len(s.items))
	for item := range s.items {
		values = append(values, item)
	}
	return values
}

// Adds the items in set Collection to desired set
func (s *Set[T]) AddAll(collection Set[T]) {
	for value := range collection.items {
		s.Add(value)
	}
}
