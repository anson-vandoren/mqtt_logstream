# MQTT forwarder for Cribl LogStream

This is a very simplistic utility to push MQTT payloads from a local (behind a firewall)
broker up to Cribl LogStream for further shaping/dissemination. It's written to send to
a LogStream Cloud `tcp_json` source over TLS.

# Usage

## Prerequisites

- [Git](https://git-scm.com/)
- [Go](https://golang.org/) (probably at least Go 1.11, tested on 1.13.8)
- A [Cribl Logstream Cloud](https://cribl.io/logstream-cloud/) account (free version is fine)
- A MQTT server you can subscribe to - if you don't already have this you're likely not interested in this project...

## Clone and build

```shell
$ git clone https://github.com/anson-vandoren/mqtt_logstream.git
$ cd mqtt_logstream
$ go build .
```

## Configure

```shell
$ copy sample_config.yml config.yml
```

Open `config.yml` and make changes appropriate to your situation, using the comments as a guide

## Run

```shell
$ ./mqtt_logstream
```