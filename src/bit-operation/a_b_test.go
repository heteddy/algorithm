package bit_operation

import "testing"

func TestAdd(t *testing.T) {
	var a uint64 = 900
	var b uint64 = 1023
	sum := Add(a, b)

	if sum == (a + b) {
		t.Log("success")
	} else {
		t.FailNow()
	}
}
