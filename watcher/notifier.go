package watcher

import "time"

// Notifier is a simple trigger model
type Notifier interface {
	Trigger() <-chan interface{}
	Cleanup()
}

func NewNotifier(trigger <-chan interface{}) Notifier {
	return NewNotifierWithCleanup(trigger, nil)
}

func NewNotifierWithCleanup(trigger <-chan interface{}, cleanup func()) Notifier {
	return &notifier{
		trigger: trigger,
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
	trigger <-chan interface{}
	cleanup func()
}

func (n *notifier) Trigger() <-chan interface{} {
	return n.trigger
}

func (n *notifier) Cleanup() {
	if n.cleanup != nil {
		n.cleanup()
	}
}
