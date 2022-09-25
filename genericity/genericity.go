package genericity

import (
	"fmt"
	"go/types"
	"strconv"
	"time"
)

// 泛型的函数
func HelloGenericity[T int](a []T) {
	for _, v := range a {
		fmt.Println(v)
	}
}

// 泛型的map
type M[K string, V any] map[K]V

// 泛型的管道
type C[T any] chan T

//泛型的类型
type Student[T any] []T

// 这样进行泛型的限制
type NumStr interface {
	Num | Str
}

// ~ 波浪号是指底层类型
type Num interface {
	~int | ~uint64 | ~uint8 | ~uint16 | ~uint32 | ~uint | ~uintptr
}

type Str interface {
	~string
}

// 通过interface中的方法来约束泛型的类型
type Price int
type Price2 string

type ShowPrice interface {
	String() string
	~int | ~string
}

func (i Price2) String() string {
	return string(i)
}

func (i Price) String() string {
	return strconv.Itoa(int(i))
}

// 以下两个有什么不同呢
func showPriceList[T ShowPrice](s []T) (ret []string) {
	for i := range s {
		ret = append(ret, s[i].String())
	}
	return
}

//func showPriceList2(s []ShowPrice) (ret []string) {
//	for i := range s {
//		ret = append(ret, s[i].String())
//	}
//	return
//}

// comparable的学习

func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

// 仿照java的泛型应用
type Nums[T comparable] []T

// 通用的sort
func (n Nums[T]) Sort() {

}

type Cloneable[T types.Array] interface {
	Clone() T
}

type ArrayCommon[T string] []T

func (a ArrayCommon[T]) Clone() ArrayCommon[T] {
	return nil
}

type IntDate int

type ITime interface {
	ToTime() time.Time
	ToDDMMYYYY() string
	ToYYYYMMDD() string
	ToYYYY_MM_DD() string
	ToDD_MM_YYYY() string
	~int | ~string
}

func (i IntDate) ToTime() *time.Time {
	return nil
}

func (i IntDate) ToDDMMYYYY() string {
	return ""
}
func (i IntDate) ToYYYYMMDD() string {
	return ""
}
func (i IntDate) ToYYYY_MM_DD() string {
	return ""
}
func (i IntDate) ToDD_MM_YYYY() string {
	return ""
}
