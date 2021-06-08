package consumer

import "github.com/alexhowarth/go-tilt"

type Message struct {
	Tilt tilt.Tilt
}

type Consumer interface {
	Listen()
	Channel() chan Message
	Configure(interface{})
}

var ConsumerTypesMap = map[string]Consumer{
	"mqtt": NewMQTT(),
	"noop": NewNOOP(),
}
