package util

import "strconv"

// GENERICS!!
func All[T comparable](a []T, same T) bool {
	for _, v := range a {
		if v != same {
			return false
		}
	}

	return true
}

func Map[T, U any](a []T, f func(s T) U) []U {
	result := []U{}
	for _, v := range a {
		new := f(v)
		result = append(result, new)
	}

	return result
}

func Atois(a []string) []int {
	return Map(a, func(x string) int {
		i, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		return i
	})
}
