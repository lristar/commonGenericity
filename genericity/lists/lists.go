package lists

import (
	"container/list"
	"myproject/go_base/base/genericity/array"
	"sync"
)

type ListType interface {
	comparable
	// 后续添加
}

// ValueStu 统一的传参结构体
type ValueStu[T ListType, V any] struct {
	Key   T
	Value V
}

// ListX 支持快速查找的双向链表 以空间换时间
type ListX[T ListType, V any] struct {
	l    *list.List
	data map[T][]*list.Element
	lmx  *sync.RWMutex
	dmx  *sync.RWMutex
}

func NewListX[T ListType, V any]() *ListX[T, V] {
	return &ListX[T, V]{
		l:    list.New(),
		data: make(map[T][]*list.Element),
		lmx:  &sync.RWMutex{},
		dmx:  &sync.RWMutex{},
	}
}

func (l *ListX[T, V]) AppendHead(k *ValueStu[T, V]) {
	e := l.inListFront(k)
	l.inMap(k.Key, e)
}

func (l *ListX[T, V]) AppendTail(k *ValueStu[T, V]) {
	e := l.inListTail(k)
	l.inMap(k.Key, e)
}

// AppendBefore
func (l *ListX[T, V]) AppendBefore(k *ValueStu[T, V], ele *list.Element) {
	e := l.inOneBefore(k, ele)
	l.inMap(k.Key, e)
}

func (l *ListX[T, V]) AppendAfter(k *ValueStu[T, V], ele *list.Element) {
	e := l.inOneAfter(k, ele)
	l.inMap(k.Key, e)
}

// 删除的操作
func (l *ListX[T, V]) PopHead() {
	if e := l.outHead(); e != nil {
		k := e.Value.(*ValueStu[T, V])
		l.outMap(k.Key, e)
	}
}

func (l *ListX[T, V]) PopTail() {
	if e := l.outTail(); e != nil {
		k := e.Value.(*ValueStu[T, V])
		l.outMap(k.Key, e)
	}
}

// todo 如何删除map中的数组中的一个元素
func (l *ListX[T, V]) PopByKey(k *ValueStu[T, V]) {
	// 需要获取
	if k == nil {
		return
	}
	src := l.getMapElement(k.Key)
	for i := range src {
		if v := src[i].Value.(*ValueStu[T, V]); v == k {
			l.outMap(k.Key, src[i])
		}
	}
}

func (l *ListX[T, V]) GetHeadElement() *list.Element {
	l.lmx.RLock()
	defer l.lmx.RUnlock()
	return l.getHead()
}

func (l *ListX[T, V]) GetTailElement() *list.Element {
	l.lmx.RLock()
	defer l.lmx.RUnlock()
	return l.getTail()
}

func (l *ListX[T, V]) GetHeadValue() *ValueStu[T, V] {
	if v := l.GetHeadElement(); v != nil {
		return v.Value.(*ValueStu[T, V])
	}
	return nil
}

func (l *ListX[T, V]) GetTailValue() *ValueStu[T, V] {
	if v := l.GetTailElement(); v != nil {
		return v.Value.(*ValueStu[T, V])
	}
	return nil
}

func (l *ListX[T, V]) GetValuesByKey(key T) []*ValueStu[T, V] {
	return l.getMap(key)
}

// 插入到双向链表的内置方法
func (l *ListX[T, V]) inListFront(k *ValueStu[T, V]) *list.Element {
	l.lmx.Lock()
	defer l.lmx.Unlock()
	return l.l.PushFront(k)
}

func (l *ListX[T, V]) inListTail(k *ValueStu[T, V]) *list.Element {
	l.lmx.Lock()
	defer l.lmx.Unlock()
	return l.l.PushBack(k)
}

func (l *ListX[T, V]) inOneBefore(k *ValueStu[T, V], node *list.Element) *list.Element {
	l.lmx.Lock()
	defer l.lmx.Unlock()
	return l.l.InsertBefore(k, node)
}

func (l *ListX[T, V]) inOneAfter(k *ValueStu[T, V], node *list.Element) *list.Element {
	l.lmx.Lock()
	defer l.lmx.Unlock()
	return l.l.InsertAfter(k, node)
}

func (l *ListX[T, V]) getHead() *list.Element {
	return l.l.Front()
}

func (l *ListX[T, V]) getTail() *list.Element {
	return l.l.Back()
}

func (l *ListX[T, V]) outTail() *list.Element {
	l.lmx.Lock()
	defer l.lmx.Unlock()
	if e := l.getTail(); e != nil {
		l.l.Remove(e)
		return e
	}
	return nil
}

func (l *ListX[T, V]) outHead() *list.Element {
	l.lmx.Lock()
	defer l.lmx.Unlock()
	if e := l.getHead(); e != nil {
		l.l.Remove(e)
		return e
	}
	return nil
}

// 插入到map的内置方法
func (l *ListX[T, V]) inMap(key T, e *list.Element) {
	l.dmx.Lock()
	defer l.dmx.Unlock()
	if ll, ok := l.data[key]; ok {
		ll = array.Append[*list.Element](ll, e)
		l.data[key] = ll
	} else {
		l.data[key] = []*list.Element{e}
	}
}

func (l *ListX[T, V]) outMap(key T, e *list.Element) {
	l.dmx.Lock()
	defer l.dmx.Unlock()
	if v, ok := l.data[key]; ok {
		v = array.Delete[*list.Element](v, e)
		l.data[key] = v
	}
}

func (l *ListX[T, V]) getMap(key T) (res []*ValueStu[T, V]) {
	l.dmx.RLock()
	defer l.dmx.RUnlock()
	if v, ok := l.data[key]; ok {
		for i := range v {
			aa := v[i].Value.(*ValueStu[T, V])
			res = array.Append[*ValueStu[T, V]](res, aa)
		}
	}
	return
}

func (l *ListX[T, V]) getMapElement(key T) (res []*list.Element) {
	l.dmx.RLock()
	defer l.dmx.RUnlock()
	if v, ok := l.data[key]; ok {
		for i := range v {
			aa := v[i]
			res = append(res, aa)
		}
	}
	return
}
