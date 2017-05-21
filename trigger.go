package chanUtils

import (
	"context"
	"sync/atomic"
)

// Trigger wakes up a waiting goroutine many times
// The number that Wait() or <-trigger is called is insured.
// Note that Wait() or <-trigger can be used by only one goroutine.
type Trigger struct {
	trigger SimpleTrigger
	counter int32
}

// Wake resume a goroutine calling Wait()
func (trigger *Trigger) Wake() {
	atomic.AddInt32(&trigger.counter, 1)
	trigger.trigger.Wake()
}

// WaitWithContext waits until trigger.Wake() or ctx.Done()
func (trigger *Trigger) WaitWithContext(ctx context.Context) error {
	for {
		if atomic.LoadInt32(&trigger.counter) > 0 {
			atomic.AddInt32(&trigger.counter, -1)
			return nil
		}
		select {
		case <-trigger.trigger:
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// Wait waits until Wake() is called.
func (trigger *Trigger) Wait() {
	trigger.WaitWithContext(context.Background())
}

// You can write like the following example
/*
var trigger Trigger
go func(){
	time.Sleep(5 * time.Second)
	trigger.Wake()
	time.Sleep(5 * time.Second)
	trigger.Wake()
}

trigger.Wait()

trigger.WaitWithContext(context.Background())
*/

// NewTrigger creates a new Trigger
func NewTrigger() *Trigger {
	return &Trigger{
		trigger: NewSimpleTrigger(),
		counter: 0,
	}
}
