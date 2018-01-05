package set

import (
	"crypto/sha1"
	"fmt"
	"log"
	"strings"
	"testing"
)

// TestAdd verifies that the set.Add() method is working properly
func TestAdd(t *testing.T) {
	log.Println("TestAdd()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results for adding new elements
	var tests = []struct {
		element interface{}
		result  bool
	}{
		// New items
		{2, true},
		{4, true},
		{6, true},
		// Existing items
		{1, false},
		{3, false},
		{5, false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		tempSet := set.Clone()

		// Attempt to add an element to the set, verify result
		ok := set.Add(test.element)
		if ok != test.result {
			t.Fatalf("set.Add(%d) - unexpected result: %t", test.element, ok)
		}

		log.Println(tempSet, "+", test.element, "=", set)
	}
}

// TestCartesianProduct verifies that the set.CartesianProduct() method is working properly
func TestCartesianProduct(t *testing.T) {
	log.Println("TestCartesianProduct()")

	// Create a set, add some initial values
	set := New(1, 2)

	// Create a table of tests and expected results of Set cartesian products
	var tests = []struct {
		source *Set
		target *Set
	}{
		// Same items
		{New(1, 2), New(Pair{1, 1}, Pair{1, 2}, Pair{2, 1}, Pair{2, 2})},
		// New items
		{New(3, 4), New(Pair{1, 3}, Pair{1, 4}, Pair{2, 3}, Pair{2, 4})},
		// Combination of items
		{New(1, 2, 3), New(Pair{1, 1}, Pair{1, 2}, Pair{1, 3}, Pair{2, 1}, Pair{2, 2}, Pair{2, 3})},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to get the cartesian product between sets, verify result
		product := set.CartesianProduct(test.source)
		if !product.Equal(test.target) {
			t.Fatalf("set.CartesianProduct() - sets not equal: %s != %s", product.String(), test.target.String())
		}

		log.Println(set, "×", test.source, "=", product)
	}
}

// TestClone verifies that the set.Clone() method is working properly
func TestClone(t *testing.T) {
	log.Println("TestClone()")

	// Create a table of tests and expected results of cloning
	var tests = []struct {
		source *Set
		target *Set
		result bool
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5), true},
		// Re-ordered items
		{New(2, 4, 6), New(6, 4, 2), true},
		// Different items
		{New(1, 2, 3), New(1, 2, 3, 4), false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to clone the current set, verify result
		if clone := test.source.Clone(); clone.Equal(test.target) != test.result {
			t.Fatalf("set.Clone() - unexpected result: %t", test.result)
		}

		log.Println(test.source)
	}
}

// TestDifference verifies that the set.Difference() method is working properly
func TestDifference(t *testing.T) {
	log.Println("TestDifference()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results of Set differences
	var tests = []struct {
		source *Set
		target *Set
	}{
		// Same items
		{New(1, 3, 5), New()},
		// New items (no difference)
		{New(2, 4, 6), New(1, 3, 5)},
		// Combination of items
		{New(1, 2, 6), New(3, 5)},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		difference := set.Difference(test.source)
		if !difference.Equal(test.target) {
			t.Fatalf("set.Difference() - sets not equal: %s != %s", difference.String(), test.target.String())
		}

		log.Println(set, "\\", test.source, "=", difference)
	}
}

// TestEnumerate verifies that the set.Enumerate() method is working properly
func TestEnumerate(t *testing.T) {
	log.Println("TestEnumerate()")

	// Create a slice of expected values upon set enumeration
	expected := []int{1, 3, 5, 7, 9}

	// Create a set
	set := New()

	// Add initial values
	for _, e := range expected {
		set.Add(e)
	}

	// Enumerate the values in the set
	for _, v := range set.Enumerate() {
		found := false

		// Check that the expected value was found upon set enumeration
		for _, e := range expected {
			if v == e {
				found = true
			}
		}

		// If value not found, test fails
		if !found {
			t.Fatalf("set.Enumerate() - element missing: %v", v)
		}
	}

	log.Println(set, "->", set.Enumerate())
}

// TestEqual verifies that the set.Equal() method is working properly
func TestEqual(t *testing.T) {
	log.Println("TestEqual()")

	// Create a table of tests and expected results of cloning
	var tests = []struct {
		source *Set
		target *Set
		result bool
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5), true},
		// Re-ordered items
		{New(2, 4, 6), New(6, 4, 2), true},
		// Repeated items
		{New(1, 2, 3), New(1, 2, 3, 1, 2), true},
		// Different items
		{New(1, 2, 3), New(1, 2, 4), false},
		// Different lengths
		{New(2, 4, 6), New(2, 4, 6, 8), false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Check set equality
		if test.source.Equal(test.target) != test.result {
			t.Fatalf("set.Equal() - unexpected result: %t", test.result)
		}

		log.Println(test.source)
	}
}

// TestFilter verifies that the set.Filter() method is working properly
func TestFilter(t *testing.T) {
	log.Println("TestFilter()")

	// Create a table of tests and expected results of Set filtering functions
	var tests = []struct {
		source *Set
		target *Set
		fn     func(interface{}) bool
	}{
		// Even number function
		{
			New(1, 2, 3, 4, 5, 6),
			New(2, 4, 6),
			func(value interface{}) bool {
				return value.(int)%2 == 0
			},
		},
		// Name filtering function
		{
			New("C", "C++", "C#", "Go", "PHP", "Ruby"),
			New("Go", "PHP", "Ruby"),
			func(value interface{}) bool {
				return !strings.HasPrefix(value.(string), "C")
			},
		},
		// Filter data type function
		{
			New(0, 1, false, true, "0", "1"),
			New(false, true),
			func(value interface{}) bool {
				_, ok := value.(bool)
				return ok
			},
		},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to apply function to set, verify result
		filterSet := test.source.Filter(test.fn)
		if !filterSet.Equal(test.target) {
			t.Fatalf("set.Filter() - sets not equal: %s != %s", filterSet.String(), test.target.String())
		}

		log.Println("filter(", test.source, ") ->", filterSet)
	}
}

// TestHas verifies that the set.Has() method is working properly
func TestHas(t *testing.T) {
	log.Println("TestHas()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results for checking membership of elements
	var tests = []struct {
		element interface{}
		result  bool
	}{
		// Existing items
		{1, true},
		{3, true},
		{5, true},
		// Non-existant items
		{2, false},
		{4, false},
		{6, false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to check if the element is contained in the set, verify result
		if ok := set.Has(test.element); ok != test.result {
			t.Fatalf("set.Has(%d) - unexpected result: %t", test.element, ok)
		}

		log.Println(test.element, "∈", set, ":", test.result)
	}
}

// TestIntersection verifies that the set.Intersection() method is working properly
func TestIntersection(t *testing.T) {
	log.Println("TestIntersection()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results of Set intersections
	var tests = []struct {
		source *Set
		target *Set
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5)},
		// New items (no intersection)
		{New(2, 4, 6), New()},
		// Combination of items
		{New(1, 2, 6), New(1)},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		intersection := set.Intersection(test.source)
		if !intersection.Equal(test.target) {
			t.Fatalf("set.Intersection() - sets not equal: %s != %s", intersection.String(), test.target.String())
		}

		log.Println(set, "∩", test.source, "=", intersection)
	}
}

// TestMap verifies that the set.Map() method is working properly
func TestMap(t *testing.T) {
	log.Println("TestMap()")

	// Create a table of tests and expected results of Set mapping functions
	var tests = []struct {
		source *Set
		target *Set
		fn     func(interface{}) interface{}
	}{
		// Square function
		{
			New(1, 3, 5),
			New(1, 9, 25),
			func(value interface{}) interface{} {
				return value.(int) * value.(int)
			},
		},
		// String replace
		{
			New("cat", "dog", "cow"),
			New("cat", "dog"),
			func(value interface{}) interface{} {
				return strings.Replace(value.(string), "cow", "cat", -1)
			},
		},
		// SHA1
		{
			New("hello", "world"),
			New("aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d", "7c211433f02071597741e6ff5a8ea34789abbf43"),
			func(value interface{}) interface{} {
				sha := sha1.New()
				sha.Write([]byte(value.(string)))
				return fmt.Sprintf("%x", sha.Sum(nil))
			},
		},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to apply function to set, verify result
		mapSet := test.source.Map(test.fn)
		if !mapSet.Equal(test.target) {
			t.Fatalf("set.Map() - sets not equal: %s != %s", mapSet.String(), test.target.String())
		}

		log.Println("map(", test.source, ") ->", mapSet)
	}
}

// TestPowerSet verifies that the set.PowerSet() method is working properly
func TestPowerSet(t *testing.T) {
	log.Println("TestPowerSet()")

	// Create a set, add some initial values
	startSet := New(1, 3, 5)
	powerSet := startSet.PowerSet()

	// Create a table of expected output from the power set
	var tests = []*Set{
		New(),
		New(1),
		New(1, 3),
		New(1, 5),
		New(1, 3, 5),
		New(3),
		New(3, 5),
		New(5),
	}

	// Iterate all sets in the power set
	for _, s := range powerSet.Enumerate() {
		set := s.(*Set)
		found := false

		// Check each set against the tests to verify it exists
		for _, test := range tests {
			if set.Equal(test) {
				found = true
				break
			}
		}

		// If set not found, test fails
		if !found {
			t.Fatalf("set.PowerSet() - set not found: %s", set.String())
		}
	}

	log.Println("P(", startSet, ") ->", powerSet)
}

// TestReduce verifies that the set.Reduce() method is working properly
func TestReduce(t *testing.T) {
	log.Println("TestReduce()")

	// Create a table of tests and expected results of Set reducing functions
	var tests = []struct {
		source *Set
		value  interface{}
		result interface{}
		fn     func(interface{}, interface{}) interface{}
	}{
		// Summing function
		{
			New(1, 2, 3),
			0,
			6,
			func(previous interface{}, value interface{}) interface{} {
				return previous.(int) + value.(int)
			},
		},
		// String transformation and concatenation
		{
			New("abc", "def", "ghi"),
			"",
			"ABCDEFGHI",
			func(previous interface{}, value interface{}) interface{} {
				return previous.(string) + strings.ToUpper(value.(string))
			},
		},
		// Accumulate and square all integers
		{
			New(2, 4, 8, false, 16, "2", "4", 32),
			0,
			1364,
			func(previous interface{}, value interface{}) interface{} {
				if _, ok := value.(int); !ok {
					return previous.(int)
				}

				return previous.(int) + (value.(int) * (value.(int)))
			},
		},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to apply function to set, verify result
		out := test.source.Reduce(test.value, test.fn)
		if out != test.result {
			t.Fatalf("set.Reduce() - unexpected result: %v", out)
		}

		log.Println("reduce(", test.source, ") ->", out)
	}
}

// TestRemove verifies that the set.Remove() method is working properly
func TestRemove(t *testing.T) {
	log.Println("TestRemove()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results for removing elements
	var tests = []struct {
		element interface{}
		result  bool
	}{
		// Existing items
		{1, true},
		{3, true},
		{5, true},
		// Non-existant items
		{2, false},
		{4, false},
		{6, false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		tempSet := set.Clone()

		// Attempt to remove an element from the set, verify result
		if ok := set.Remove(test.element); ok != test.result {
			t.Fatalf("set.Remove(%d) - unexpected result: %t", test.element, ok)
		}

		log.Println(tempSet, "-", test.element, "=", set)
	}
}

// TestSize verifies that the set.Size() method is working properly
func TestSize(t *testing.T) {
	log.Println("TestSize()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Verify initial size
	if set.Size() != 3 {
		t.Fatalf("set.Size() - unexpected result: %d", set.Size())
	}

	// Create a table of tests and expected size when adding new elements
	var tests = []struct {
		element interface{}
		size    int
	}{
		// New items
		{2, 4},
		{4, 5},
		{6, 6},
		// Existing items
		{1, 6},
		{3, 6},
		{5, 6},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Add an element to the set, check size
		set.Add(test.element)

		if set.Size() != test.size {
			t.Fatalf("set.Size()- unexpected result: %d", set.Size())
		}

		log.Println("|", set, "| =", set.Size())
	}

}

// TestSubset verifies that the set.Subset() method is working properly
func TestSubset(t *testing.T) {
	log.Println("TestSubset()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results of Set subsets
	var tests = []struct {
		source *Set
		result bool
	}{
		// Empty set
		{New(), true},
		// Same items
		{New(1, 3, 5), true},
		// New items
		{New(2, 4, 6), false},
		// Combination of items
		{New(1, 2, 6), false},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to check if set is a subset, verify result
		if set.Subset(test.source) != test.result {
			t.Fatalf("set.Subset() - unexpected result: %t", test.result)
		}

		log.Println(set, "⊆", test.source, ":", test.result)
	}
}

// TestSymmetricDifference verifies that the set.SymmetricDifference() method is working properly
func TestSymmetricDifference(t *testing.T) {
	log.Println("TestSymmetricDifference()")

	// Create a table of tests and expected results of Set symmetricDifferences
	var tests = []struct {
		source *Set
		target *Set
		result *Set
	}{
		// Same items
		{New(1, 3, 5), New(1, 2, 3), New(2, 5)},
		// Different items
		{New(2, 4, 6), New(1, 3, 5), New(1, 2, 3, 4, 5, 6)},
		// Combination of items
		{New(1, 2, 6), New(1, 5), New(2, 5, 6)},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		symDiff := test.source.SymmetricDifference(test.target)
		if !symDiff.Equal(test.result) {
			t.Fatalf("set.SymmetricDifference() - sets not equal: %s != %s", symDiff.String(), test.result.String())
		}

		log.Println(test.source, "∆", test.target, "=", symDiff)
	}
}

// TestUnion verifies that the set.Union() method is working properly
func TestUnion(t *testing.T) {
	log.Println("TestUnion()")

	// Create a set, add some initial values
	set := New(1, 3, 5)

	// Create a table of tests and expected results of Set unions
	var tests = []struct {
		source *Set
		target *Set
	}{
		// Same items
		{New(1, 3, 5), New(1, 3, 5)},
		// New items
		{New(2, 4, 6), New(1, 2, 3, 4, 5, 6)},
		// Combination of items
		{New(1, 2, 3), New(1, 2, 3, 5)},
	}

	// Iterate test table, checking results
	for _, test := range tests {
		// Attempt to add an element to the set, verify result
		union := set.Union(test.source)
		if !union.Equal(test.target) {
			t.Fatalf("set.Union() - sets not equal: %s != %s", union.String(), test.target.String())
		}

		log.Println(set, "∪", test.source, "=", union)
	}
}
