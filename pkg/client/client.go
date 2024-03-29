package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Options   *mqtt.ClientOptions
	client    mqtt.Client
	brokerUrl string
	clientId  string
	username  string
	password  string
}

type option func(*Client) error

func WithBrokerUrl(brokerUrl string) option {
	return func(c *Client) error {
		c.brokerUrl = brokerUrl
		return nil
	}
}

func WithClientId(clientId string) option {
	return func(c *Client) error {
		c.clientId = clientId
		return nil
	}
}

func WithUserInfo(username, password string) option {
	return func(c *Client) error {
		c.username = username
		c.password = password
		return nil
	}
}

func NewClient(options ...option) (*Client, error) {
	c := &Client{}

	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	clientOptions := mqtt.NewClientOptions().
		AddBroker(c.brokerUrl).
		SetClientID(c.clientId).
		SetUsername(c.username).
		SetPassword(c.password)

	clientOptions.OnConnect = defaultConnectHandler
	clientOptions.OnConnectionLost = defaultConnectLostHandler

	c.Options = clientOptions
	return c, nil
}

func (c *Client) Connect() error {
	c.client = mqtt.NewClient(c.Options)
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to connect: %s", token.Error().Error())
	}
	return nil
}

func (c *Client) Disconnect() {
	c.client.Disconnect(250)
}

func (c *Client) Subscribe(topic string, handler mqtt.MessageHandler) error {
	if token := c.client.Subscribe(topic, byte(0), handler); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe: %s", token.Error().Error())
	}
	return nil
}

func (c *Client) Unsubscribe(topic string) error {
	if token := c.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to unsubscribe: %s", token.Error().Error())
	}
	return nil
}

func (c *Client) Publish(topic string, payload any) error {
	if token := c.client.Publish(topic, byte(0), false, payload); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to publish: %s", token.Error().Error())
	}
	return nil
}

func defaultConnectHandler(client mqtt.Client) {
	log.Info("connected to mqtt")
}

func defaultConnectLostHandler(client mqtt.Client, err error) {
	log.Infof("connection lost: %v", err)
}
