package chanUtils

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestInitTrigger(t *testing.T) {
	trigger := NewTrigger()

	if trigger == nil {
		t.Error("trigger mustn' be nil")
	}
}

func TestTrigger(t *testing.T) {
	trigger := NewTrigger()
	var status int32

	atomic.StoreInt32(&status, 0)
	go func() {
		trigger.Wait()
		atomic.StoreInt32(&status, 1)
	}()

	trigger.Wake()

	time.Sleep(1 * time.Second)

	if atomic.LoadInt32(&status) != 1 {
		t.Error("Wake is not worked.")
	}

}
