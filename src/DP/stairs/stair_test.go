package stairs

import (
	"testing"
	"time"
)

func TestStairsWithoutCache(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02 15:04:05.000000"))
	t.Log(StairsWithoutCache(40))
	t.Log(time.Now().Format("2006-01-02 15:04:05.000000"))
	t.Log(SequenceStairCounter(40))
	t.Log(time.Now().Format("2006-01-02 15:04:05.000000"))
	t.Log(StairCounter(40))
	t.Log(time.Now().Format("2006-01-02 15:04:05.000000"))
}