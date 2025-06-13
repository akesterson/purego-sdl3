package mem

import "unsafe"

// Copy copies the values of a C array into a Go slice.
// If the array is nil, this function returns nil.
func Copy[T any, N ~int32](array *T, count N) []T {
	if array == nil {
		return nil
	}
	result := make([]T, count)
	copy(result, unsafe.Slice(array, count))
	return result
}

func DeepCopy[T any, N ~int32](array **T, count N) []*T {
	if array == nil {
		return nil
	}
	dst := make([]*T, count)
	src := unsafe.Slice(array, count)
	for i := 0; i < int(count); i++ {
		if src[i] != nil {
			value := *src[i]
			dst[i] = &value
		}
	}
	return dst
}
