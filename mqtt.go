package main

import (
	"log"
	"time"

	"ansonvandoren.com/mqtt_logstream/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttListener(configuration config.BrokerConfig, handler mqtt.MessageHandler) mqtt.Client {
	// set some reasonable defaults for the MQTT client
	opts := mqtt.NewClientOptions().AddBroker(configuration.Address).SetClientID(configuration.ClientID)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	// configure the message handler that will be used for all subscriptions
	opts.SetDefaultPublishHandler(handler)

	// create the client
	c := mqtt.NewClient(opts)

	// connect to the MQTT broker and make sure the connection succeeds
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe to the configured topics, using the default handler
	for _, topic := range configuration.Topics {
		if token := c.Subscribe(topic+"/#", 0, nil); token.Wait() && token.Error() != nil {
			// try to subscribe to any other topics, and don't fail
			log.Println(token.Error())
		}
	}

	return c
}
