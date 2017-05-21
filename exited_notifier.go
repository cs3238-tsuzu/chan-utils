package chanUtils

import (
	"context"
	"sync/atomic"
)

// ExitedNotifier When you call Finish(), all waiting goroutines will be resumed.
type ExitedNotifier struct {
	Channel chan bool
	exited  int32
}

// Finish resumes all waiting goroutines
func (en ExitedNotifier) Finish() {
	if atomic.CompareAndSwapInt32(&en.exited, int32(0), int32(1)) {
		close(en.Channel)
	}
}

// Wait waits Finish() is called.
func (en ExitedNotifier) Wait() {
	<-en.Channel
}

// WaitWithContext waits until Finish() is called or ctx is triggered.
func (en ExitedNotifier) WaitWithContext(ctx context.Context) {
	select {
	case <-ctx.Done():
	case <-en.Channel:
	}
}

// TriggerOrCancel sets f() to be called when Finish() is called.
// The returned function is a canceller.
// Note that you must call this when you no longer need f() to be called.
func (en ExitedNotifier) TriggerOrCancel(f func()) func() {
	canceler := make(chan bool)
	go func() {
		select {
		case <-en.Channel:
			f()
		case <-canceler:
			return
		}
	}()

	return func() {
		close(canceler)
	}
}

// NewExitedNotifier creates a new ExitedNotifier
func NewExitedNotifier() ExitedNotifier {
	return ExitedNotifier{
		Channel: make(chan bool),
		exited:  0,
	}
}
