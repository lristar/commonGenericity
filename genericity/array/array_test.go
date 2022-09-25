package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	a := ArraysCp[float64]{6.14, 1.325, 3.32, 7.21, 8.98}
	a.Sort()
	fmt.Println(a)
}

func TestAppend(t *testing.T) {
	a := Append[string]([]string{}, []string{"1", "2", "4", "3"}...)
	fmt.Println(a)
	fmt.Println(Delete(a, "4"))
}

func TestArraysCp_Delete(t *testing.T) {
	a := Append[string]([]string{}, []string{"1", "2", "4", "3"}...)
	fmt.Println(Delete(a, "5"))

	// 空数组是否会报错
	no := Append[string]([]string{}, []string{}...)
	fmt.Println(Delete(no, "2"))
}
