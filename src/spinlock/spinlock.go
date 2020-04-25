/*
@Copyright:
*/
/*
@Time : 2020/3/24 14:24
@Author : teddy
@File : spinlock.go
*/

package spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type SpinLock struct {
	routineID uint64
	count     int
	_lock     uint32
}

func (sl *SpinLock) Lock() {
	self := GoroutineID()
	if sl.routineID == self { // 如果当前线程已经获取到了锁，线程数增加一，然后返回
		sl.count++
		return
	}
	// 如果没获取到锁，则通过CAS自旋
	for !atomic.CompareAndSwapUint32((*uint32)(&sl._lock), 0, 1) {
		runtime.Gosched()
	}
}
func (sl *SpinLock) Unlock() {
	if sl.routineID != GoroutineID() {
		panic("非自己则报错")
	}
	if sl.count > 0 { // 如果大于0，表示当前线程多次获取了该锁，释放锁通过count减一来模拟
		sl.count--
	} else { // 如果count==0，可以将锁释放，这样就能保证获取锁的次数与释放锁的次数是一致的了。
		atomic.StoreUint32((*uint32)(&sl._lock), 0)
	}
}

func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}
