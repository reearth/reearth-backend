package util

func TryMap[T any, V any](collection []T, iteratee func(v T, i int) (V, error)) ([]V, error) {
	m := make([]V, 0, len(collection))
	for i, e := range collection {
		j, err := iteratee(e, i)
		if err != nil {
			return nil, err
		}
		m = append(m, j)
	}
	return m, nil
}

// https://github.com/samber/lo/issues/54
func All[T any](collection []T, predicate func(T) bool) bool {
	for _, e := range collection {
		if !predicate(e) {
			return false
		}
	}
	return true
}

// https://github.com/samber/lo/issues/54
func Any[T any](collection []T, predicate func(T) bool) bool {
	for _, e := range collection {
		if predicate(e) {
			return true
		}
	}
	return false
}
