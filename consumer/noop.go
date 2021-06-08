package consumer

import (
	"log"
	"time"
)

type NOOP struct {
	channel   chan Message
	listening bool
}

func NewNOOP() *NOOP {
	return &NOOP{
		channel: make(chan Message),
	}
}

func (n *NOOP) Configure(interface{}) {}

func (n *NOOP) Channel() chan Message {
	return n.channel
}

func (n *NOOP) Listen() {
	if n.listening {
		log.Printf("Already listening. Return\n")
		return
	}
	go func() {
		for {
			select {
			case message := <-n.channel:
				log.Printf("Consumed: %+v\n", message)
			case <-time.After(time.Millisecond):
			}
		}
	}() // pass m?
	n.listening = true
}
