/**
  * @author wellsjiang
  * @date 2018/8/14
  */

package utils

import "time"

type Timer struct {
	start time.Time
	last  time.Time
}

func (t *Timer) Start() {
	t.Reset()
}

func (t *Timer) Reset() {
	t.start = time.Now()
	t.last = t.start
}

// 单次计时，返回ns
func (t *Timer) Elapsed() time.Duration {
	elapsed := time.Since(t.last)
	t.last = time.Now()
	return t.NsToMs(elapsed)
}

// 总耗时，返回ms
func (t *Timer) TotalElapsed() time.Duration {
	return t.NsToMs(time.Since(t.start))
}

func (t *Timer) NsToMs(elapsed time.Duration) time.Duration {
	return elapsed / time.Millisecond
}

func (t *Timer) NsToS(elapsed time.Duration) time.Duration {
	return elapsed / time.Millisecond / time.Microsecond
}
