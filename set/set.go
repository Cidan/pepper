package set

import (
	"fmt"
	"sync"
)

// Set represents an unordered collection of unique values
type Set struct {
	// Mutex to allow safe, concurrent access
	mutex sync.RWMutex
	// Empty struct consumes no memory, so we just use the map keys
	m map[interface{}]struct{}
}

// New creates a new Set, and initializes its internal map, optionally adding initial elements to the set
func New(values ...interface{}) *Set {
	// Initialize set
	s := Set{
		m: make(map[interface{}]struct{}),
	}

	// If items are specified in the initializer, immediately add them to the set
	for _, v := range values {
		s.Add(v)
	}

	return &s
}

// Add inserts a new element into the set, returning true if the element was newly added, or false
// if it already existed
func (s *Set) Add(value interface{}) bool {
	// Check existence
	found := s.Has(value)

	// Lock set for write
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Add value to set
	s.m[value] = struct{}{}

	// Return inverse, so true if element already existed
	return !found
}

// Pair represents a pair of elements created from a cartesian product
type Pair struct {
	X interface{}
	Y interface{}
}

// String returns a string representation of this pair
func (p *Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

// CartesianProduct returns a set containing ordered pairs of every permutation between two sets
func (s *Set) CartesianProduct(t *Set) *Set {
	// Create a set of ordered pair permutations between the sets
	cpSet := New()

	// Enumerate the source set
	for _, x := range s.Enumerate() {
		// Enumerate the target set
		for _, y := range t.Enumerate() {
			// Create pair, insert elements, insert into set
			cpSet.Add(Pair{
				X: x,
				Y: y,
			})
		}
	}

	return cpSet
}

// Clone copies the current set into a new, identical set
func (s *Set) Clone() *Set {
	// Copy set into a new set
	outSet := New()
	for _, v := range s.Enumerate() {
		outSet.Add(v)
	}

	return outSet
}

// Difference returns a set containing all elements present in this set, but without any elements
// present in the parameter set
func (s *Set) Difference(t *Set) *Set {
	// Create a set of differences between the sets
	diffSet := New()

	// Enumerate and check all elements in the current set
	for _, e := range s.Enumerate() {
		found := false

		// Check if element is present in parameter set
		for _, p := range t.Enumerate() {
			// Element found
			if e == p {
				found = true
			}
		}

		// If element was not found, add it to diff set
		if !found {
			diffSet.Add(e)
		}
	}

	return diffSet
}

// Enumerate returns an unordered slice of all elements in the set
func (s *Set) Enumerate() []interface{} {
	// Lock set for read
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Gather all values into a slice
	values := make([]interface{}, 0)
	for k := range s.m {
		values = append(values, k)
	}

	return values
}

// Equal returns whether or not two sets have the same length and no differences, meaning they are equal
func (s *Set) Equal(t *Set) bool {
	return s.Size() == t.Size() && s.Difference(t).Size() == 0
}

// Filter applies a function over all elements of the set, and returns all elements which return true
// when the function is applied
func (s *Set) Filter(fn func(interface{}) bool) *Set {
	// Create a set to return with elements which match filter function
	filterSet := New()

	// Enumerate all elements and apply the function
	for _, e := range s.Enumerate() {
		// Apply the function, add elements which it matches
		if fn(e) {
			filterSet.Add(e)
		}
	}

	return filterSet
}

// Has checks for membership of an element in the set
func (s *Set) Has(value interface{}) bool {
	// Lock set for read
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Check for value
	if _, ok := s.m[value]; ok {
		// Value found
		return true
	}

	// Value not found
	return false
}

// Intersection returns a set containing all elements present in both the current set and the parameter set
func (s *Set) Intersection(t *Set) *Set {
	// Copy current set, create a set of intersections between the sets
	intSet := s.Clone()

	// Get all differences between the sets
	for _, d := range s.Difference(t).Enumerate() {
		// Remove all different elements
		intSet.Remove(d)
	}

	return intSet
}

// Map applies a function over all elements of the set, and returns the resulting set
func (s *Set) Map(fn func(interface{}) interface{}) *Set {
	// Create a set to return with function applied
	mapSet := New()

	// Enumerate all elements and apply the function
	for _, e := range s.Enumerate() {
		// Apply the function, capture result
		mapSet.Add(fn(e))
	}

	return mapSet
}

// powerSet is called recursively to generate the power set, or set containing all
// possible subsets
func powerSet(set *Set) *Set {
	// Create the output set
	pSet := New()

	// If set is empty, return the nil set
	if set.Size() == 0 {
		pSet.Add(New())
		return pSet
	}

	// Get the head element of the set, and remove it
	head := set.Enumerate()[0]
	set.Remove(head)

	// Get the "tail" power set, and add it to the set
	tpSet := powerSet(set)
	for _, tp := range tpSet.Enumerate() {
		pSet.Add(tp)
	}

	// Iterate the "tail" power set and get the "head" power set,
	// and add it to the set
	for _, tp := range tpSet.Enumerate() {
		hSet := tp.(*Set).Clone()
		hSet.Add(head)
		pSet.Add(hSet)
	}

	return pSet
}

// PowerSet generates a set of all possible subsets, given the current set
func (s *Set) PowerSet() *Set {
	// If set is empty, return set of empty set
	if s.Size() == 0 {
		pSet := New()
		pSet.Add(New())
		return pSet
	}

	// Call the recursive powerSet function on a clone of this set
	return powerSet(s.Clone())
}

// Reduce applies a function over all elements of the set, accumulating the results into a final result value
func (s *Set) Reduce(value interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	// Enumerate all elements and apply the function
	for _, e := range s.Enumerate() {
		value = fn(value, e)
	}

	return value
}

// Remove destroys an element in the set, returning true if the element was destroyed, or false if it did not exist
func (s *Set) Remove(value interface{}) bool {
	// Check existence
	found := s.Has(value)

	// Lock set for write
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Remove value from set
	delete(s.m, value)

	return found
}

// Size returns the size or cardinality of this set
func (s *Set) Size() int {
	// Lock set for read
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.m)
}

// String returns a string representation of this set
func (s *Set) String() string {
	// Print identifier
	str := "{ "

	// Check for empty set, print symbol if empty
	if s.Size() == 0 {
		return str + "Ø }"
	}

	// Print all elements
	for k := range s.m {
		// Print pairs separately
		if pair, ok := k.(Pair); ok {
			str = str + fmt.Sprintf("%v ", pair.String())
		} else {
			str = str + fmt.Sprintf("%v ", k)
		}
	}

	return str + "}"
}

// Subset determines if a parameter set is a subset of elements within this set, returning true if it
// is a subset, or false if it is not
func (s *Set) Subset(t *Set) bool {
	// Check if all elements in the parameter set are contained within the set
	for _, v := range t.Enumerate() {
		// Check if element is contained, if not, return false
		if !s.Has(v) {
			return false
		}
	}

	return true
}

// SymmetricDifference returns a set containing all elements which are not shared between this set
// and the parameter set
func (s *Set) SymmetricDifference(t *Set) *Set {
	return s.Difference(t).Union(t.Difference(s))
}

// Union returns a set containing all elements present in this set, as well as all elements present
// in the parameter set
func (s *Set) Union(t *Set) *Set {
	// Clone the current set into a new set
	outSet := s.Clone()

	// Enumerate and add all elements from the parameter set
	for _, e := range t.Enumerate() {
		outSet.Add(e)
	}

	return outSet
}
