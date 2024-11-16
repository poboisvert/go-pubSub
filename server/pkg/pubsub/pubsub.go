package pubsub

import "sync"

// PubSub manages subscribers and notifications.
type PubSub struct {
	sync.Mutex
	Subscribers map[int]*Subscriber
}

// NewPubSub creates a new PubSub instance.
func NewPubSub() *PubSub {
	return &PubSub{
		Subscribers: make(map[int]*Subscriber),
	}
}

// AddSubscriber adds a new subscriber.
func (ps *PubSub) AddSubscriber(sub *Subscriber) {
	ps.Lock()
	defer ps.Unlock()
	ps.Subscribers[sub.ID] = sub
}

// RemoveSubscriber removes a subscriber.
func (ps *PubSub) RemoveSubscriber(subID int) {
	ps.Lock()
	defer ps.Unlock()
	delete(ps.Subscribers, subID)
}

// Notify sends a message to all subscribers.
func (ps *PubSub) Notify(message string) {
	ps.Lock()
	defer ps.Unlock()
	for _, sub := range ps.Subscribers {
		sub.Chan <- message
	}
}
