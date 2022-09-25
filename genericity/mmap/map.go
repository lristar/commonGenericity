package mmap

import (
	"github.com/lristar/commonGenericity/genericity/array"
	"golang.org/x/exp/constraints"
	"sync"
)

type Set[T comparable] map[T]interface{}

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

func (s Set[T]) Del(k T) {
	delete(s, k)
}

type HSet[T constraints.Ordered] struct {
	m   map[T]interface{}
	n   array.ArraysCp[T]
	len int
	mx  sync.RWMutex
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
