package main

import (
	"encoding/json"
	"fmt"

	"github.com/cvancleave/go-mqtt/pkg/client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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
		panic(err)
	}

	// optionally set other client options here to override defaults

	// connect
	if err := c.Connect(); err != nil {
		panic(err)
	}

	// subscribe with below handler
	if err := c.Subscribe(topic, handler); err != nil {
		panic(err)
	}

	fmt.Printf("subscribed to topic: %s\n", topic)

	for range make(chan string) {
	}
}

func handler(c mqtt.Client, m mqtt.Message) {
	var data map[string]any
	if err := json.Unmarshal(m.Payload(), &data); err != nil {
		panic(err)
	}
	fmt.Printf("message received from %s: %v\n", m.Topic(), data["text"])
}
