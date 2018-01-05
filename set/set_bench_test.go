package set

import (
	"testing"
)

// BenchmarkAdd checks the performance of the set.Add() method
func BenchmarkAdd(b *testing.B) {
	// Create a new set
	set := New()

	// Run set.Add() b.N times
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

// benchmarkCartesianProduct checks the performance of the set.CartesianProduct() method
func benchmarkCartesianProduct(n int, s *Set, t *Set) {
	// Run set.CartesianProduct() n times
	for i := 0; i < n; i++ {
		s.CartesianProduct(t)
	}
}

// BenchmarkCartesianProductSmall checks the performance of the set.CartesianProduct() method
// over a small data set
func BenchmarkCartesianProductSmall(b *testing.B) {
	benchmarkCartesianProduct(b.N, New(1, 2), New(2, 1))
}

// BenchmarkCartesianProductLarge checks the performance of the set.CartesianProduct() method
// over a large data set
func BenchmarkCartesianProductLarge(b *testing.B) {
	benchmarkCartesianProduct(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkClone checks the performance of the set.Clone() method
func benchmarkClone(n int, s *Set) {
	// Run set.Clone() n times
	for i := 0; i < n; i++ {
		s.Clone()
	}
}

// BenchmarkCloneSmall checks the performance of the set.Clone() method
// over a small data set
func BenchmarkCloneSmall(b *testing.B) {
	benchmarkClone(b.N, New(1, 2))
}

// BenchmarkCloneLarge checks the performance of the set.Clone() method
// over a large data set
func BenchmarkCloneLarge(b *testing.B) {
	benchmarkClone(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// benchmarkDifference checks the performance of the set.Difference() method
func benchmarkDifference(n int, s *Set, t *Set) {
	// Run set.Difference() n times
	for i := 0; i < n; i++ {
		s.Difference(t)
	}
}

// BenchmarkDifferenceSmall checks the performance of the set.Difference() method
// over a small data set
func BenchmarkDifferenceSmall(b *testing.B) {
	benchmarkDifference(b.N, New(1, 2), New(2, 1))
}

// BenchmarkDifferenceLarge checks the performance of the set.Difference() method
// over a large data set
func BenchmarkDifferenceLarge(b *testing.B) {
	benchmarkDifference(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkEnumerate checks the performance of the set.Enumerate() method
func benchmarkEnumerate(n int, s *Set) {
	// Run set.Enumerate() n times
	for i := 0; i < n; i++ {
		s.Enumerate()
	}
}

// BenchmarkEnumerateSmall checks the performance of the set.Enumerate() method
// over a small data set
func BenchmarkEnumerateSmall(b *testing.B) {
	benchmarkEnumerate(b.N, New(1, 2))
}

// BenchmarkEnumerateLarge checks the performance of the set.Enumerate() method
// over a large data set
func BenchmarkEnumerateLarge(b *testing.B) {
	benchmarkEnumerate(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// benchmarkEqual checks the performance of the set.Equal() method
func benchmarkEqual(n int, s *Set, t *Set) {
	// Run set.Equal() n times
	for i := 0; i < n; i++ {
		s.Equal(t)
	}
}

// BenchmarkEqualSmall checks the performance of the set.Equal() method
// over a small data set
func BenchmarkEqualSmall(b *testing.B) {
	benchmarkEqual(b.N, New(1, 2), New(2, 1))
}

// BenchmarkEqualLarge checks the performance of the set.Equal() method
// over a large data set
func BenchmarkEqualLarge(b *testing.B) {
	benchmarkEqual(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkFilter checks the performance of the set.Filter() method
func benchmarkFilter(n int, s *Set, fn func(interface{}) bool) {
	// Run set.Filter() n times
	for i := 0; i < n; i++ {
		s.Filter(fn)
	}
}

// BenchmarkFilterSmall checks the performance of the set.Filter() method
// over a small data set
func BenchmarkFilterSmall(b *testing.B) {
	benchmarkFilter(b.N, New(1, 2), func(v interface{}) bool {
		return v.(int)%2 == 0
	})
}

// BenchmarkFilterLarge checks the performance of the set.Filter() method
// over a large data set
func BenchmarkFilterLarge(b *testing.B) {
	benchmarkFilter(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), func(v interface{}) bool {
		return v.(int)%2 == 0
	})
}

// BenchmarkHas checks the performance of the set.Has() method
func BenchmarkHas(b *testing.B) {
	// Create a new set
	set := New()

	// Run set.Has() b.N times
	for i := 0; i < b.N; i++ {
		set.Has(i)
	}
}

// benchmarkIntersection checks the performance of the set.Intersection() method
func benchmarkIntersection(n int, s *Set, t *Set) {
	// Run set.Intersection() n times
	for i := 0; i < n; i++ {
		s.Intersection(t)
	}
}

// BenchmarkIntersectionSmall checks the performance of the set.Intersection() method
// over a small data set
func BenchmarkIntersectionSmall(b *testing.B) {
	benchmarkIntersection(b.N, New(1, 2), New(2, 1))
}

// BenchmarkIntersectionLarge checks the performance of the set.Intersection() method
// over a large data set
func BenchmarkIntersectionLarge(b *testing.B) {
	benchmarkIntersection(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkMap checks the performance of the set.Map() method
func benchmarkMap(n int, s *Set, fn func(interface{}) interface{}) {
	// Run set.Map() n times
	for i := 0; i < n; i++ {
		s.Map(fn)
	}
}

// BenchmarkMapSmall checks the performance of the set.Map() method
// over a small data set
func BenchmarkMapSmall(b *testing.B) {
	benchmarkMap(b.N, New(1, 2), func(v interface{}) interface{} {
		return v.(int) * v.(int)
	})
}

// BenchmarkMapLarge checks the performance of the set.Map() method
// over a large data set
func BenchmarkMapLarge(b *testing.B) {
	benchmarkMap(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), func(v interface{}) interface{} {
		return v.(int) * v.(int)
	})
}

// benchmarkPowerSet checks the performance of the set.PowerSet() method
func benchmarkPowerSet(n int, s *Set) {
	// Run set.PowerSet() n times
	for i := 0; i < n; i++ {
		s.PowerSet()
	}
}

// BenchmarkPowerSetSmall checks the performance of the set.PowerSet() method
// over a small data set
func BenchmarkPowerSetSmall(b *testing.B) {
	benchmarkPowerSet(b.N, New(1, 2))
}

// BenchmarkPowerSetLarge checks the performance of the set.PowerSet() method
// over a large data set
func BenchmarkPowerSetLarge(b *testing.B) {
	benchmarkPowerSet(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

// benchmarkReduce checks the performance of the set.Reduce() method
func benchmarkReduce(n int, s *Set, fn func(interface{}, interface{}) interface{}) {
	// Run set.Reduce() n times
	for i := 0; i < n; i++ {
		s.Reduce(i, fn)
	}
}

// BenchmarkReduceSmall checks the performance of the set.Reduce() method
// over a small data set
func BenchmarkReduceSmall(b *testing.B) {
	benchmarkReduce(b.N, New(1, 2), func(p interface{}, v interface{}) interface{} {
		return p.(int) + v.(int)
	})
}

// BenchmarkReduceLarge checks the performance of the set.Reduce() method
// over a large data set
func BenchmarkReduceLarge(b *testing.B) {
	benchmarkReduce(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), func(p interface{}, v interface{}) interface{} {
		return p.(int) + v.(int)
	})
}

// BenchmarkRemove checks the performance of the set.Remove() method
func BenchmarkRemove(b *testing.B) {
	// Create a new set
	set := New()

	// Run set.Remove() b.N times
	for i := 0; i < b.N; i++ {
		set.Remove(i)
	}
}

// BenchmarkSize checks the performance of the set.Size() method
func BenchmarkSize(b *testing.B) {
	// Create a new set
	set := New()

	// Run set.Size() b.N times
	for i := 0; i < b.N; i++ {
		set.Size()
	}
}

// benchmarkSubset checks the performance of the set.Subset() method
func benchmarkSubset(n int, s *Set, t *Set) {
	// Run set.Subset() n times
	for i := 0; i < n; i++ {
		s.Subset(t)
	}
}

// BenchmarkSubsetSmall checks the performance of the set.Subset() method
// over a small data set
func BenchmarkSubsetSmall(b *testing.B) {
	benchmarkSubset(b.N, New(1, 2), New(2, 1))
}

// BenchmarkSubsetLarge checks the performance of the set.Subset() method
// over a large data set
func BenchmarkSubsetLarge(b *testing.B) {
	benchmarkSubset(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkSymmetricDifference checks the performance of the set.SymmetricDifference() method
func benchmarkSymmetricDifference(n int, s *Set, t *Set) {
	// Run set.SymmetricDifference() n times
	for i := 0; i < n; i++ {
		s.SymmetricDifference(t)
	}
}

// BenchmarkSymmetricDifferenceSmall checks the performance of the set.SymmetricDifference() method
// over a small data set
func BenchmarkSymmetricDifferenceSmall(b *testing.B) {
	benchmarkSymmetricDifference(b.N, New(1, 2), New(2, 1))
}

// BenchmarkSymmetricDifferenceLarge checks the performance of the set.SymmetricDifference() method
// over a large data set
func BenchmarkSymmetricDifferenceLarge(b *testing.B) {
	benchmarkSymmetricDifference(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}

// benchmarkUnion checks the performance of the set.Union() method
func benchmarkUnion(n int, s *Set, t *Set) {
	// Run set.Union() n times
	for i := 0; i < n; i++ {
		s.Union(t)
	}
}

// BenchmarkUnionSmall checks the performance of the set.Union() method
// over a small data set
func BenchmarkUnionSmall(b *testing.B) {
	benchmarkUnion(b.N, New(1, 2), New(2, 1))
}

// BenchmarkUnionLarge checks the performance of the set.Union() method
// over a large data set
func BenchmarkUnionLarge(b *testing.B) {
	benchmarkUnion(b.N, New(1, 2, 3, 4, 5, 6, 7, 8, 9), New(9, 8, 7, 6, 5, 4, 3, 2, 1))
}
