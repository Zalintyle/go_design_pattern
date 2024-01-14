package singleton_pattern

import (
	"sync"
	"sync/atomic"
)

type MyOnce struct {
	done uint32
	m    sync.Mutex
}

func (o *MyOnce) Do(f func()) {
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
