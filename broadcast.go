package broadcast

import "sync"

// Broadcast allows to send a signal to all listeners
type Broadcast struct {
	lock sync.RWMutex
	ch   chan struct{}
}

// NewBroadcast creates a new broadcast
func NewBroadcast() *Broadcast {
	return &Broadcast{
		lock: sync.RWMutex{},
		ch:   make(chan struct{}),
	}
}

// Receive a channel on which the next (close) signal will be sent
func (b *Broadcast) Receive() <-chan struct{} {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.ch
}

// Send a signal to all listeners
func (b *Broadcast) Send() {
	b.lock.Lock()
	defer b.lock.Unlock()
	close(b.ch)
	b.ch = make(chan struct{})
}
