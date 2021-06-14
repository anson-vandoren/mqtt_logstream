package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ansonvandoren.com/mqtt_logstream/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type LogstreamConnection struct {
	address string
	conn    *tls.Conn
	encoder *json.Encoder
}

// NewLogstreamConnection generates a `LogstreamConnection` based on provided config, connects
// to Logstream, and sends any header fields if present in the config
func NewLogstreamConnection(config config.LogstreamConfig) *LogstreamConnection {
	c := new(LogstreamConnection)

	host := config.Address
	port := config.Port
	address := fmt.Sprintf("%s:%d", host, port)
	c.address = address

	// open TLS/TCP connection to LogStream
	conn, err := tls.Dial("tcp", c.address, nil)
	if err != nil {
		log.Panicf("error dialing %s: %s\n", c.address, err.Error())
	}
	log.Printf("opened TCP conn to %s\n", c.address)
	c.conn = conn
	c.encoder = json.NewEncoder(c.conn)

	// possibly send some header data to the stream
	authToken := config.AuthToken
	fields := config.Fields
	header := LogstreamHeader{authToken, fields}
	if fields != nil || authToken != "" {
		err = c.encoder.Encode(&header)
		printableHeader, _ := json.Marshal(header)
		log.Printf("sent header to LogStream: %s\n", printableHeader)
		if err != nil {
			log.Panicf("failed to encode JSON for %v: %s\n", header, err)
		}
	}

	return c
}

func (c *LogstreamConnection) Close() {
	c.conn.Close()
}

type LogstreamJson struct {
	Topic     string
	Payload   string `json:"_raw"`
	Timestamp int64  `json:"_time"`
}

type LogstreamHeader struct {
	AuthToken string            `json:"authToken"`
	Fields    map[string]string `json:"fields"`
}

func (lc *LogstreamConnection) Handler(_ mqtt.Client, message mqtt.Message) {
	msg := LogstreamJson{
		Topic:     message.Topic(),
		Payload:   string(message.Payload()),
		Timestamp: time.Now().Unix(),
	}
	log.Printf("sending %v\n", msg)
	err := lc.encoder.Encode(msg)
	if err != nil {
		fmt.Printf("error sending message: %s with topic: %s to %s\n", message.Payload(), message.Topic(), lc.address)
	}
}
