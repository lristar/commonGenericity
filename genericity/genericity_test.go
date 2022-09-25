package genericity

import (
	"testing"
)

func TestHelloGenericity(t *testing.T) {
	HelloGenericity[int]([]int{1, 3, 423, 23})
}

func TestStudent_Get(t *testing.T) {
	s := Student[int]{1, 2, 5, 23, 56, 42}
	s.Get()
}
