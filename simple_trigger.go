package chanUtils

import (
	"context"
)

// SimpleTrigger wakes up one waiting goroutine many times.
// The number that Wait() or <-trigger is called is NOT insured.
// Note that Wait() or <-trigger can be used by only one goroutine.
type SimpleTrigger chan bool

// Wake resume a goroutine calling Wait()
func (trigger SimpleTrigger) Wake() {
	select {
	case trigger <- true:
	default:
	}
}

// Wait waits until Wake() is called.
func (trigger SimpleTrigger) Wait() {
	<-trigger
}

// WaitWithContext waits until Wake() or ctx.Done()
func (trigger SimpleTrigger) WaitWithContext(ctx context.Context) error {
	select {
	case <-trigger:
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// You can write like the following example
/*
var trigger SimpleTrigger
go func(){
	time.Sleep(5 * time.Second)
	trigger.Wake()
	time.Sleep(5 * time.Second)
	trigger.Wake()
}

<-trigger

trigger.Wait()
*/

// NewSimpleTrigger creates a new Trigger
func NewSimpleTrigger() SimpleTrigger {
	return make(chan bool, 1)
}
