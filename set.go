package set

import (
	"encoding/json"
	"fmt"
)

// A generic set. You can create with
//    s := make(Set[int])
//    s = make(Set[string]) /// etc
//    s = MakeSet([]int{1,2,3,4,5})
type Set[T comparable] map[T]struct{}

func FromSlice[T comparable](slice []T) Set[T] {
	s := make(Set[T], len(slice))
	for _, v := range slice {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Contains(v T) bool {
	_, found := (s)[v]
	return found
}

func (s Set[T]) ToSlice() []T {
	keys := make([]T, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
func (s Set[T]) Length() int {
	return len(s)
}

// Add set to this set. This is the same as Set.Union, except that
// it is an inplace operation
func (a *Set[T]) AddSet(b Set[T]) {
	for k := range b {
		(*a)[k] = struct{}{}
	}
}

// Add individual value(s) to the set.
func (s *Set[T]) Add(vals ...T) {
	for _, v := range vals {
		(*s)[v] = struct{}{}
	}

}

// Adds the records from the slice.
func (s *Set[T]) AddSlice(slice []T) {
	for _, v := range slice {
		(*s)[v] = struct{}{}
	}
}

// Returns the difference a - b That is to say:
//   a := MakeSet([]int{1,2,3,4,5})
//   b := MakeSet([]int{1,2,3})
//   d := a.Difference(b)
// d contains 4,5
func (a Set[T]) Difference(b Set[T]) Set[T] {
	diff := make(Set[T], len(a))
	for v := range a {
		if _, found := b[v]; !found {
			diff[v] = struct{}{}
		}
	}
	return diff
}

// Returns the union of set a + b. This is the equivilent of set.Add, but returns
// a new set instead of being an inplace operation.
func (s Set[T]) Union(b Set[T]) Set[T] {
	for k := range b {
		s[k] = struct{}{}
	}
	return s
}

// Returns the intersection of sets a and b, that is to say only things that are in both 'a'
// and 'b'
func (a Set[T]) Intersection(b Set[T]) Set[T] {
	intersection := make(Set[T], len(a))
	for v := range a {
		if _, found := b[v]; found {
			intersection[v] = struct{}{}
		}
	}
	return intersection
}

func (s Set[T]) String() string {
	return fmt.Sprint(s.ToSlice())
}

// A set marshalls into a slice.
func (s Set[T]) MarshalJSON() ([]byte, error) {
	list := s.ToSlice()
	return json.Marshal(list)
}

// A set unmarshalls from a slice.
func (s *Set[T]) UnmarshalJSON(b []byte) error {
	var list []T
	err := json.Unmarshal(b, &list)
	if err != nil {
		return err
	}
	(*s) = FromSlice(list)
	return nil
}

func (s Set[T]) MarshalText() ([]byte, error) {
	return s.MarshalJSON()
}

func (s *Set[T]) UnmarshalText(text []byte) error {
	return s.UnmarshalJSON(text)
}
