package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/cvancleave/go-mqtt/pkg/client"
)

func main() {

	// set info
	topic := "my/custom/topic"
	brokerUrl := "tcp://localhost:1883"
	clientId := "clientB"
	username := ""
	password := ""

	// create client
	c, err := client.NewClient(client.WithInfo(brokerUrl, clientId, username, password))
	if err != nil {
		log.Fatal(err)
	}

	// optionally set other client options here to override defaults

	// connect
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	// setup payload
	payload, err := json.Marshal("hi mom")
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("publishing payload")

	// publish
	if err := c.Publish(topic, payload); err != nil {
		log.Fatal(err)
	}

	for range make(chan string) {
	}
}