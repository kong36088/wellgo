/**
  * @author wellsjiang
  * @date 2018/8/14
  */

package utils

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := &Timer{}

	timer.Start()

	time.Sleep(2 * time.Second)
	t.Logf("wellgo: elapsed=%dms", timer.Elapsed())
	t.Logf("wellgo: total elapsed=%dms", timer.TotalElapsed())
}
