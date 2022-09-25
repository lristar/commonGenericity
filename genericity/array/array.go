package array

import (
	"golang.org/x/exp/constraints"
	"sort"
)

type CMP interface {
	constraints.Ordered
}

type (
	ArraysCp[T CMP] []T
)

func (a ArraysCp[T]) Sort() {
	sort.Sort(&a)
}

func (a ArraysCp[T]) Len() int {
	return len(a)
}

func (a ArraysCp[T]) Less(i, j int) bool {
	return a[i] < a[j]
}

// Swap swaps the elements with indexes i and j.
func (a ArraysCp[T]) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ArraysCp[T]) GetMax() T {
	a.Sort()
	return a[len(a)-1]
}

func (a ArraysCp[T]) GetMin() T {
	a.Sort()
	return a[0]
}

func (a ArraysCp[T]) Mid() (T, int) {
	a.Sort()
	ll := len(a)
	if ll%2 == 0 {
		return a[ll/2] + a[ll/2+1], 2
	}
	return a[ll/2+1], 1
}

func (a ArraysCp[T]) Sum() (sum T) {
	for i := range a {
		sum += a[i]
	}
	return
}

func Delete[T comparable](src []T, v T) []T {
	delId := 0
	if len(src) == 0 {
		return src
	}
	for i := range src {
		if src[i] == v {
			delId = i
			break
		}
		if i == len(src)-1 {
			return src
		}
	}
	return Append[T]([]T{}, Append(src[:delId], src[delId+1:]...)...)
}

func Append[T comparable](src []T, v ...T) []T {
	c := make([]T, len(src)+len(v))
	copy(c, src)
	for i := range v {
		c[len(src)+i] = v[i]
	}
	return c
}

// 后边是任意的数组
