package main

import "fmt"

type void struct{}

type Set[T comparable] struct {
	data map[T]void
}

// КонструкторЫ*
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]void)}
}
func NewSetFromSlice[T comparable](items []T) *Set[T] {
	s := NewSet[T]()
	for _, v := range items {
		s.Add(v)
	}
	return s

}
func (s *Set[T]) Add(v T) {
	s.data[v] = void{}
}
func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}
func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

//Пересечение текущего мн с другим
func (s *Set[T]) intersection(set *Set[T]) *Set[T] {
	resultSet := NewSet[T]()
	for k := range s.data {
		if set.Has(k) {
			resultSet.Add(k)
		}
	}
	return resultSet
}

func main() {
	set1 := NewSetFromSlice[int]([]int{1, 2, 3})
	set2 := NewSetFromSlice[int]([]int{1, 3})
	set3 := set1.intersection(set2)
	for k := range set3.data {
		fmt.Println(k)
	}
}
