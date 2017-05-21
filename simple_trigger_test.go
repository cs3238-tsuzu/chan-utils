package chanUtils

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestInitSimpleTrigger(t *testing.T) {
	trigger := NewSimpleTrigger()

	if trigger == nil {
		t.Error("simpleTrigger mustn' be nil")
	}
}

func TestSimpleTrigger(t *testing.T) {
	trigger := NewSimpleTrigger()
	var status int32

	atomic.StoreInt32(&status, 0)
	go func() {
		trigger.Wait()
		atomic.StoreInt32(&status, 1)
	}()

	trigger.Wake()

	time.Sleep(1 * time.Second)

	if atomic.LoadInt32(&status) != 1 {
		t.Error("Wake do not worked.")
	}

}
