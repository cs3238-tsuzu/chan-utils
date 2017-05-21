package chanUtils

import (
	"context"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestInitTrigger(t *testing.T) {
	trigger := NewTrigger()

	if trigger.trigger == nil || trigger.counter != 0 {
		t.Error("trigger must be not nil and counter must be 0")
	}
}

func TestTrigger(t *testing.T) {
	func() {
		trigger := NewTrigger()
		var status int32

		atomic.StoreInt32(&status, 0)
		done := make(chan bool)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				trigger.Wait()
				atomic.AddInt32(&status, 1)

				select {
				case <-done:
					return
				default:
				}
			}
		}()

		time.Sleep(1 * time.Second)

		cnt := int32(rand.Intn(500))
		for i := int32(0); i < cnt; i++ {
			trigger.Wake()
		}
		time.Sleep(1 * time.Second)
		close(done)

		if atomic.LoadInt32(&status) != cnt {
			t.Error("Triger did not worked.")
			t.FailNow()
		}

		trigger.Wake()
		wg.Wait()
	}()

	func() {
		trigger := NewTrigger()
		var status int32

		atomic.StoreInt32(&status, 0)
		done := make(chan bool)
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				ctx, f := context.WithTimeout(context.Background(), 300*time.Millisecond)
				defer f()
				err := trigger.WaitWithContext(ctx)
				if err == nil {
					atomic.AddInt32(&status, 1)
				}

				select {
				case <-done:
					return
				default:
				}
			}
		}()

		time.Sleep(1 * time.Second)

		cnt := int32(rand.Intn(5))
		for i := int32(0); i < cnt; i++ {
			trigger.Wake()
			time.Sleep(1 * time.Second)
		}
		close(done)

		wg.Wait()
		if atomic.LoadInt32(&status) != cnt {
			t.Error("Triger did not worked.")
			t.FailNow()
		}
	}()
}
