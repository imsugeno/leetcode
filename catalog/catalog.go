package catalog

import "cmp"

func chmin[T cmp.Ordered](a *T, b T) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax[T cmp.Ordered](a *T, b T) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func abs[T signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
