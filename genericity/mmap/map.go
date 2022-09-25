package mmap

import (
	"golang.org/x/exp/constraints"
	"myproject/go_base/base/genericity/array"
	"sort"
	"sync"
)

type Set[T string | int] map[T]interface{}

type HSet[T constraints.Ordered] struct {
	m   map[T]interface{}
	n   array.ArraysCp[T]
	len int
	mx  sync.RWMutex
}

func (s Set[T]) Set(k T) {
	s[k] = struct{}{}
}

func (s Set[T]) Len() int {
	return s.Len()
}

func (s Set[T]) IsExist(k T) bool {
	if _, ok := s[k]; ok {
		return true
	}
	return false
}

func (a HSet[T]) Add(key T) {
	if _, ok := a.m[key]; !ok {
		a.mx.Lock()
		defer a.mx.Unlock()
		a.m[key] = struct{}{}
		a.len++
		c := make(array.ArraysCp[T], a.len)
		copy(c, a.n)
		c[a.len-1] = key
		c.Sort()
		a.n = c
	}
}

func (a HSet[T]) Del(key T) {
	if _, ok := a.m[key]; ok {
		a.mx.Lock()
		defer a.mx.Unlock()
		delete(a.m, key)
		a.len--

	}
}

func (a HSet[T]) Sort() {
	sort.Sort(a)
}
