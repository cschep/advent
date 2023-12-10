package util

// GENERICS!!
func All[T comparable](a []T, same T) bool {
	for _, v := range a {
		if v != same {
			return false
		}
	}

	return true
}
