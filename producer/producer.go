// Package producer consists of a Publisher and Consumer(s)
package producer

import (
	"log"
	"reflect"
	"time"

	"github.com/alexhowarth/go-ferment-monitor/consumer"
)

type Publisher struct {
	Channel   chan consumer.Message
	Consumers []consumer.Consumer
	started   bool
}

func NewPublisher() *Publisher {
	return &Publisher{
		Channel: make(chan consumer.Message),
	}
}

func (p *Publisher) AddConsumer(c consumer.Consumer) {
	if p.started {
		log.Printf("Consumers already started - rejecting add\n")
		return
	}
	p.Consumers = append(p.Consumers, c)
}

func (p *Publisher) StartConsumers() {
	if p.started {
		log.Printf("Consumers already started\n")
		return
	}
	p.started = true
	for _, c := range p.Consumers {
		log.Printf("Starting consumer\n")
		c.Listen()
	}

	go func() {
		for {
			select {
			case message := <-p.Channel:
				for _, subscriber := range p.Consumers {
					log.Printf("Publishing to: %v: %+v", reflect.TypeOf(subscriber).String(), message)
					subscriber.Channel() <- message
				}
			case <-time.After(time.Millisecond):
			}
		}
	}()
}
