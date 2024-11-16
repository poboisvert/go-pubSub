package pubsub

// Subscriber holds subscription details.
type Subscriber struct {
	ID   int
	Chan chan string
}

// NewSubscriber creates a new subscriber instance.
func NewSubscriber(id int) *Subscriber {
	return &Subscriber{
		ID:   id,
		Chan: make(chan string),
	}
}
