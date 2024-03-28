package main

import (
	"encoding/json"

	"github.com/cvancleave/go-mqtt/pkg/client"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

func main() {

	// set info
	topic := "my/custom/topic"
	brokerUrl := "tcp://localhost:1883"
	clientId := "clientA"
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

	// subscribe with below handler
	if err := c.Subscribe(topic, handler); err != nil {
		log.Fatal(err)
	}

	log.Infof("subscribed to topic: %s", topic)

	for range make(chan string) {
	}
}

func handler(c mqtt.Client, m mqtt.Message) {
	var data string
	if err := json.Unmarshal(m.Payload(), &data); err != nil {
		log.Errorf("failed to unmarshal payload: %s", err.Error())
		return
	}
	log.Infof("message received from %s: %s", m.Topic(), data)
}
