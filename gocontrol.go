package gocontrol

import (
	"sync/atomic"
)

type Guard struct {
	counter int64
}

func (guard *Guard) Go() func() {
	atomic.AddInt64(&guard.counter, 1)
	return func() {
		atomic.AddInt64(&guard.counter, -1)
	}
}

func (guard Guard) AliveN() int64 {
	return atomic.LoadInt64(&guard.counter)
}
