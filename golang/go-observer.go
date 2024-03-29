package main

import "fmt"

// Event defines an indication of a some occurence
type Event struct {
	Data int                // Data in this case is a simple int.
}

// Observer defines a standard interface to listen for a specific event.
type Observer interface {
	OnNotify(Event)         // OnNotify allows to publsh an event
}

// Notifier is the instance being observed.
type  Notifier interface {
	Register(Observer)      // Register itself to listen/observe events.
	Unregister(Observer)    // Remove itself from the collection of observers/listeners.
	Notify(Event)           // Notify publishes new events to listeners.
}

// Example

type observer struct {
	id int
}

func (o *observer) OnNotify(e Event) {
	fmt.Printf("observer %d recieved event %d\n", o.id, e.Data)
}

type notifier struct {
	observers map[Observer]struct{}
}

func (n *notifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *notifier) Unregister(o Observer) {
	delete(n.observers, o)
}

func (n *notifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}

func main() {
	n := notifier{
        observers: map[Observer]struct{}{},
	}

	n.Register(&observer{1})
	n.Register(&observer{2})

	n.Notify(Event{1})
	n.Notify(Event{101})
	n.Notify(Event{9999})
}
