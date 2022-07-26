package watcher

import "time"

// Notifier is a simple notifier model
type Notifier interface {
	Notify() <-chan interface{}
	Cleanup()
}

func NewNotifier(ch <-chan interface{}) Notifier {
	return NewNotifierWithCleanup(ch, nil)
}

func NewNotifierWithCleanup(ch <-chan interface{}, cleanup func()) Notifier {
	return &notifier{
		ch:      ch,
		cleanup: cleanup,
	}
}

func NewTimerNotifier(interval time.Duration) Notifier {
	ch := make(chan interface{})
	stopCh := make(chan struct{})
	notifier := NewNotifierWithCleanup(ch, func() { stopCh <- struct{}{} })

	t := time.NewTicker(interval)
	go func() {
		defer t.Stop()
		for {
			select {
			case v := <-t.C:
				ch <- v
			case <-stopCh:
				return
			}
		}
	}()
	return notifier
}

type notifier struct {
	ch      <-chan interface{}
	cleanup func()
}

func (n *notifier) Notify() <-chan interface{} {
	return n.ch
}

func (n *notifier) Cleanup() {
	if n.cleanup != nil {
		n.cleanup()
	}
}

// Trigger is a simple trigger model
type Trigger interface {
	Notifier

	Trigger() chan<- interface{}
}

func NewTrigger() Trigger {
	ch := make(chan interface{})
	return &trigger{
		Notifier:  NewNotifierWithCleanup(ch, nil),
		triggerCh: ch,
	}
}

type trigger struct {
	Notifier
	triggerCh chan<- interface{}
}

func (t *trigger) Trigger() chan<- interface{} {
	return t.triggerCh
}
