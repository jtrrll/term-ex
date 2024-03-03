package set

import mapset "github.com/deckarep/golang-set/v2"

// Creates a new set containing elements present in every given set
func Intersection[T comparable](sets ...mapset.Set[T]) mapset.Set[T] {
	intersection := mapset.NewSet[T]()
	for i, set := range sets {
		if i == 0 {
			intersection = set.Clone()
			continue
		}
		intersection = intersection.Intersect(set)
	}
	return intersection
}
