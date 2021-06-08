package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alexhowarth/go-ferment-monitor/models"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	channel   chan Message
	listening bool
	client    mqtt.Client
}

func NewMQTT() *MQTT {
	return &MQTT{
		channel: make(chan Message),
	}
}

func (m *MQTT) Configure(i interface{}) {
	conf := i.(models.MQTTConf)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(conf.Url)
	opts.SetUsername(conf.Username)
	opts.SetPassword(conf.Password)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Unable to connect to MQTT: %+v", token.Error())
	}

	m.client = c
}

func (m *MQTT) Channel() chan Message {
	return m.channel
}

func (m *MQTT) Listen() {
	if m.listening {
		log.Printf("Already listening. Return\n")
		return
	}
	go func() {
		for {
			select {
			case message := <-m.channel:
				log.Printf("Consumed: %+v\n", message)
				payload := struct {
					//Beer string  `json:"beer"`
					C    float64 `json:"temp_c"`
					F    float64 `json:"temp_f"` // should probably put c and f here for others to use (or can grafana do it?)
					Sg   float64 `json:"sg"`
					Col  string  `json:"col"`
					Beer string  `json:"beer"`
					Og   float64 `json:"og"`
				}{
					//Beer: "beer005",
					C:    message.Tilt.Celsius(),
					F:    message.Tilt.Fahrenheit(),
					Sg:   message.Tilt.Gravity(),
					Og:   0.005,
					Col:  message.Tilt.Colour().String(),
					Beer: "The Foo Beer",
				}
				json, err := json.Marshal(payload)
				if err != nil {
					log.Printf("Error with JSON marshalling: %v\n", err)
					return
				}
				log.Printf("Sending JSON: %v\n", string(json)) // log topic
				topic := fmt.Sprintf("tilt/%v", strings.ToLower(message.Tilt.Colour().String()))
				log.Printf("topic: %+v", topic)
				// token := m.c.Publish("tilt/purple/beer005", 0, false, json)
				token := m.client.Publish(topic, 0, false, json)
				token.Wait()
			case <-time.After(time.Millisecond):
			}
		}
	}()
	m.listening = true
}
