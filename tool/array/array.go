package array

func Map[T any, R any](arr []T, fn func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

func Unique[T comparable](arr []T) []T {
	uniqueMap := make(map[T]struct{})
	for _, v := range arr {
		uniqueMap[v] = struct{}{}
	}

	result := make([]T, 0, len(uniqueMap))
	for k := range uniqueMap {
		result = append(result, k)
	}
	return result
}

func ToMap[T any, K comparable](arr []T, fn func(T) K) map[K]T {
	result := make(map[K]T)
	for _, v := range arr {
		k := fn(v)
		result[k] = v
	}
	return result
}
