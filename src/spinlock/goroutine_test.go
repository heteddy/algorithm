/*
@Copyright:
*/
/*
@Time : 2020/3/24 14:28
@Author : teddy
@File : goroutine_test.go
*/

package spinlock

import "testing"

func TestGoroutineID(t *testing.T) {
	_id1 := GoroutineID()
	_id2 := GoroutineID()
	var _id3 uint64
	ch := make(chan struct{})
	go func() {
		_id3 = GoroutineID()
		ch <- struct{}{}
	}()
	<-ch
	if _id1 != _id2 {
		t.FailNow()
	}
	if _id1 == _id3 {
		t.FailNow()
	}
}
