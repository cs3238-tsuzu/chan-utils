package chanUtils

// Trigger wakes up a waiting goroutine many times
// Note that Wait() or <-trigger can be used by only one goroutine.
type Trigger chan bool

// Wake resume a goroutine calling Wait()
func (trigger Trigger) Wake() {
	select {
	case trigger <- true:
	default:
	}
}

// Wait watis until Wake() is called.
func (trigger Trigger) Wait() {
	<-trigger
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

<-trigger

trigger.Wait()
*/

// NewTrigger creates a new Trigger
func NewTrigger() Trigger {
	return make(chan bool, 1)
}
