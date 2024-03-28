package main

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/cvancleave/go-mqtt/pkg/client"
)

func main() {

	// set variables
	topic := "my/custom/topic"

	// create client options
	optA := client.WithBrokerUrl("tcp://localhost:1883")
	optB := client.WithClientId("clientB")

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

	// publish with below payload
	if err := c.Publish(topic, payload()); err != nil {
		log.Fatal(err)
	}

	log.Infof("published payload")
}

func payload() []byte {
	data := map[string]any{
		"time": time.Now().Unix(),
		"text": "hi mom",
	}
	payload, _ := json.Marshal(data)
	return payload
}
