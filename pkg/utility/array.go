package utility

func Contains[T comparable](array []T, elem T) bool {
	for _, e := range array {
		if e == elem {
			return true
		}
	}
	return false
}
