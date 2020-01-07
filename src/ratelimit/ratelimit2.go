package ratelimit

import (
	"sync"
	"time"
)

const infinityDuration time.Duration = 0x7fffffffffffffff

type Clock interface {
	Now() time.Time
	Sleep(time.Duration)
}

type RealClock struct {
}

func (r RealClock) Now() time.Time {
	return time.Now()
}

func (r RealClock) Sleep(d time.Duration) {
	time.Sleep(d)
}

// todo 支持预热，
// 把10ms作为一个tick
type Bucket struct {
	clock       Clock
	startTime   time.Time
	mu          sync.Mutex
	capacity    int64
	quanta      int64         // 一次放入的数量    一个间隔放入的数量
	interval    time.Duration //每次放的间隔  默认为10ms 1e7  10*time.Millisecond
	available   int64         //当前的tokens数量
	latestTicks int64         //
}
type RateLimiter interface {
	Acquire() bool
	TryAcquire() bool
	AcquireTokens(tokens int64) bool
	TryAcquireTokens(tokens int64, d time.Duration) bool
	Rate() int64
	SetRate(tokensPerSec int64)
	Available() int64
	TakeAvailable(count int64) int64
}

func (tb *Bucket) Acquire() bool {
	duration, ok := tb.takeToken(time.Now(), 1, infinityDuration);
	if ok {
		if duration > 0 {
			time.Sleep(duration)
		}
		return true
	}
	return false
}

// 获取尽可能多的token
func (tb *Bucket) TakeAvailable(count int64) int64 {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	if count <= 0 {
		return 0
	}
	tb.adjustAvailableToken(tb.currentTicks(time.Now()))
	if tb.available <= 0 {
		return 0
	}
	if count > tb.available {
		count = tb.available
	}
	tb.available -= count
	return count
}

func (tb *Bucket) Available() int64 {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.adjustAvailableToken(tb.currentTicks(time.Now()))
	return tb.available
}

func (tb *Bucket) TryAcquire() bool {
	duration, ok := tb.takeToken(time.Now(), 1, infinityDuration);
	if ok {
		if duration == 0 {
			return true
		}
	}
	return false
}

func (tb *Bucket) AcquireTokens(tokens int64) bool {
	if tokens < 0 {
		return true
	}
	duration, ok := tb.takeToken(time.Now(), tokens, infinityDuration);
	if ok {
		if duration == 0 {
			return true
		}
	}
	return true
}
func (tb *Bucket) TryAcquireTokens(tokens int64, d time.Duration) bool {
	if tokens < 0 {
		return true
	}
	duration, ok := tb.takeToken(time.Now(), tokens, d);
	if ok {
		if duration > 0 {
			time.Sleep(duration)
			return true
		}
	}
	return false
}

//每秒多少个
func (tb *Bucket) Rate() int64 {
	return int64(time.Second/tb.interval) * tb.quanta
}

func (tb *Bucket) SetRate(tokensPerSec int64) {

}

// 距离上一次获取的时间差
func (tb *Bucket) currentTicks(now time.Time) int64 {
	// startTime是启动时间
	return int64(now.Sub(tb.startTime) / tb.interval)
}

// 每次获取available或者acquire 都需要调整token
func (tb *Bucket) adjustAvailableToken(tick int64) {
	lastTick := tb.latestTicks
	tb.latestTicks = tick
	if tb.available >= tb.capacity {
		tb.available = tb.capacity
		return
	}
	tb.available += (tick - lastTick) * tb.quanta
	if tb.available >= tb.capacity {
		tb.available = tb.capacity
	}
}

func (tb *Bucket) takeToken(now time.Time, count int64, maxWait time.Duration) (time.Duration, bool) {
	if count < 0 {
		return 0, true
	}
	ticks := tb.currentTicks(now)
	tb.adjustAvailableToken(ticks)
	_avail := tb.available - count
	// 如果当前的available已经足够
	if _avail > 0 {
		return 0, true
	} else {
		// 当前的available不够用
		// 计算还需要多久的时间
		left := count - tb.available //还差left个token
		//int64(left / tb.interval)  //left token所需要的时间
		// todo: 这里是left还是
		endTick := ticks + left/tb.quanta
		endTime := tb.startTime.Add(time.Duration(endTick) * tb.interval)

		waitDuration := endTime.Sub(now)

		if waitDuration > maxWait {
			return 0, false
		}
		tb.available = _avail
		return waitDuration, true
	}
}
