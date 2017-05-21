package chanUtils

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestInitExitedNotifier(t *testing.T) {
	notifier := NewExitedNotifier()

	if notifier.Channel == nil || notifier.Exited != 0 {
		t.Error("Channel must be not nil and Exited must be 0")
	}
}

func TestTriggerOrCancel(t *testing.T) {
	func() {
		notifier := NewExitedNotifier()
		var status int32

		atomic.StoreInt32(&status, 0)
		go func() {
			notifier.TriggerOrCancel(func() {
				atomic.StoreInt32(&status, 1)
			})
		}()

		time.Sleep(1 * time.Second)

		notifier.Finish()

		time.Sleep(1 * time.Second)

		if atomic.LoadInt32(&status) != 1 {
			t.Error("Triger was not worked.")
			t.FailNow()
		}
	}()

	func() {
		notifier := NewExitedNotifier()
		var status int32

		atomic.StoreInt32(&status, 0)
		go func() {
			canceller := notifier.TriggerOrCancel(func() {
				atomic.StoreInt32(&status, 1)
			})

			canceller()
		}()

		time.Sleep(1 * time.Second)

		notifier.Finish()

		time.Sleep(1 * time.Second)

		if atomic.LoadInt32(&status) != 0 {
			t.Error("canceller was not worked.")
			t.FailNow()
		}
	}()
}

func TestWait(t *testing.T) {
	func() {
		notifier := NewExitedNotifier()
		var status int32

		atomic.StoreInt32(&status, 0)
		go func() {
			notifier.Wait()
			atomic.StoreInt32(&status, 1)
		}()

		time.Sleep(1 * time.Second)

		notifier.Finish()

		time.Sleep(1 * time.Second)

		if atomic.LoadInt32(&status) != 1 {
			t.Error("Wait was not worked.")
			t.FailNow()
		}
	}()
}

func TestWaitWithContext(t *testing.T) {
	func() {
		notifier := NewExitedNotifier()
		var status int32

		atomic.StoreInt32(&status, 0)
		ctx, f := context.WithCancel(context.Background())
		go func() {
			notifier.WaitWithContext(ctx)

			atomic.StoreInt32(&status, 1)
		}()
		f()

		time.Sleep(1 * time.Second)

		if atomic.LoadInt32(&status) != 1 {
			t.Error("WaitWithContext was not worked.")
			t.FailNow()
		}

		notifier.Finish()
	}()
}
