package singleton

import (
	"sync"
	"sync/atomic"
)

type MyOnce struct {
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/386),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	// m Mutex used to lock and unlock the done variable.
	m sync.Mutex
}

func (o *MyOnce) Do(f func()) {
	// Note: Here is an incorrect implementation of Do:
	//
	//	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	//		f()
	//	}
	//
	// Do guarantees that when it returns, f has finished.
	// This implementation would not implement that guarantee:
	// given two simultaneous calls, the winner of the cas would
	// call f, and the second would return immediately, without
	// waiting for the first's call to f to complete.
	// This is why the slow path falls back to a mutex, and why
	// the atomic.StoreUint32 must be delayed until after f returns.

	if atomic.LoadUint32(&o.done) == 0 {
		// do f()
		o.doSlow(f)
	}
	// else do nothing
}

func (o *MyOnce) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		// 这里使用defer，而不是直接执行 atomic.StoreUint32(&o.done, 1) 是因为：
		// 另一个协程可能在 o.done 的值更新完，但是更新o.done的协程执行完f()之前，直接return
		// 如果使用defer，也就是获取到锁的协程需要先执行 f()，执行完毕之后才会更新 o.done -> 1
		// 保证了不会有其他协程因为o.done==1而不等待更新o.done的协程执行完f()，就直接return
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
