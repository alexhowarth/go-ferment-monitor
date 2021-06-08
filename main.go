package main

import (
	"log"
	"runtime"
	"time"

	"github.com/alexhowarth/go-ferment-monitor/consumer"
	"github.com/alexhowarth/go-ferment-monitor/models"
	"github.com/alexhowarth/go-ferment-monitor/producer"
	"github.com/alexhowarth/go-ferment-monitor/server"
	"github.com/alexhowarth/go-tilt"
)

func main() {
	models.Init()

	p := producer.NewPublisher()
	println("wft")

	// this needs to be elsewhere, constantly checking to see if enabled has been set on any service and starting it
	// use channels to restart? reconfigure etc?

	mqttModel := &models.MQTTModel{}
	conf, err := mqttModel.GetMQTTConfig()

	if err != nil {
		log.Printf("Error loading config - Please configure by going to http://localhost:5000: %+v", err)
	}

	log.Printf("Config: %+v", conf)

	if conf.Enabled {
		if c, ok := consumer.ConsumerTypesMap["mqtt"]; ok {
			log.Printf("Adding consumer: MQTT\n")
			c.Configure(conf)
			p.AddConsumer(c)
		}
	}

	p.StartConsumers()

	tiltChan := make(chan tilt.Tilt)
	tiltConf := tilt.Config{Channel: tiltChan}

	go func() {
		tiltModel := &models.TiltModel{}
		for {
			select {
			case t := <-tiltChan:
				log.Printf("Producing to tiltChan %v", t)
				p.Channel <- consumer.Message{Tilt: t}
				conf := models.TiltConf{Colour: t.Colour()}
				tiltModel.CreateTilt(conf) // this needs to be first or create so that we can pass name and og to p.Channel <- above
			case <-time.After(time.Millisecond):
			}
		}
	}()

	t := tiltConf.NewScanner()

	scan := 120 * time.Second
	sleep := 120 * time.Second

	// TODO maybe we just populate a 'found' table here and display it as a table on the first tab

	go func() {
		for {
			log.Printf("Goroutines: %v", runtime.NumGoroutine())
			log.Printf("Scanning for %v", scan)
			t.Scan(scan)
			log.Printf("Sleeping for %v", sleep)
			time.Sleep(sleep)
		}
	}()

	server.Init()
}
