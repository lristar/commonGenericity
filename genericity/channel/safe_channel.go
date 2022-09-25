package coroutine

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

type SafeChannel[T any] struct {
	//mx       *sync.RWMutex
	one      *sync.Once
	isClosed *atomic.Value
	c        chan T
}

func NewSafeChannel[T any]() *SafeChannel[T] {
	v := atomic.Value{}
	v.Store(false)
	return &SafeChannel[T]{
		one:      &sync.Once{},
		isClosed: &v,
		c:        make(chan T),
	}
}

func (c *SafeChannel[T]) Send(ctx context.Context, val T) (err error) {
	defer func() {
		if err = recover().(error); err != nil {
			c.isClosed.Store(true)
			log.Println("channel is panic")
		}
	}()
	if c.isClosed.Load() == true {
		return
	}
	c.c <- val
	return
}

// Close 改用了once
func (c *SafeChannel[T]) Close() {
	c.one.Do(func() {
		close(c.c)
		c.isClosed.Store(true)
	})
}

func (c *SafeChannel[T]) Receive(f func(v T)) (err error) {
	for {
		select {
		case v, ok := <-c.c:
			if !ok {
				return fmt.Errorf("channel is closed")
			}
			f(v)
		}
	}
}
