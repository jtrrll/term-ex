package util

// Creates a new slice containing elements present in every given slice
func Intersection[T comparable](slices ...[]T) []T {
	// 1. Initialize an empty intersection set
	set := map[T]struct{}{}
	// 2. Process all given slices
	for i, slice := range slices {
		if i == 0 {
			// 2.1. Intialize the intersection with all elements in the first slice
			for _, v := range slice {
				set[v] = struct{}{}
			}
		} else {
			// 2.2. Remove elements from the intersection that are not in the slice
			hash := map[T]struct{}{}
			for _, v := range slice {
				hash[v] = struct{}{}
			}
			for v := range set {
				_, ok := hash[v]
				if !ok {
					delete(set, v)
				}
			}
		}
	}
	// 3. Convert the intersection set to a slice
	slice := []T{}
	for v := range set {
		slice = append(slice, v)
	}
	return slice
}
