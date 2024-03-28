package main

import (
	"encoding/json"

	"github.com/cvancleave/go-mqtt/pkg/client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

func main() {

	// set variables
	topic := "my/custom/topic"

	// create client options
	optA := client.WithBrokerUrl("tcp://localhost:1883")
	optB := client.WithClientId("clientA")

	// create client
	c, err := client.NewClient(optA, optB)
	if err != nil {
		log.Fatal(err)
	}

	// optionally set other client options here to override defaults

	// connect
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	// subscribe with below handler
	if err := c.Subscribe(topic, handler); err != nil {
		log.Fatal(err)
	}

	log.Infof("subscribed to topic: %s", topic)

	for range make(chan string) {
	}
}

func handler(c mqtt.Client, m mqtt.Message) {
	var data map[string]any
	if err := json.Unmarshal(m.Payload(), &data); err != nil {
		log.Errorf("failed to unmarshal payload: %s", err.Error())
		return
	}
	log.Infof("message received from %s: %v", m.Topic(), data["text"])
}
