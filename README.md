## What's this?
- ChanUtils is a utility library for goroutins and channels in Go
- Channels and goroutines are very useful, but some utilities are necessary to use them comfortably.

## Classes
- ExitedNotifier
    - When you call Finish(), all waiting goroutines will be resumed.
- SimpleTrigger
    - This can wakes up one waiting goroutine many times.
    - The number that Wait() or <-trigger is called is NOT insured.
- Trigger
    - This can wakes up one waiting goroutine many times.
    - The number that Wait() or WaitWithContext() is called is insured.

## License
- Under the MIT License
- Copyright (c) 2017 Tsuzu

## Examples
- Not prepared yet.
- Please refer to *_test.go