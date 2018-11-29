package main

import (
	"flag"
)

type HttpConfig struct {
	Listen string
}

type MqttConfig struct {
	Host     string
	User     string
	Password string

	BaseTopic string
}

func (config MqttConfig) BrokerUri() string {
	uri := "tcp://"
	if config.User != "" {
		uri += config.User

		if config.Password != "" {
			uri += ":" + config.Password
		}

		uri += "@"
	}

	uri += config.Host

	return uri
}

type Config struct {
	Mqtt     *MqttConfig
	Http     *HttpConfig
	RepoPath string
	GpioPin  int
}

func parseFlags() *Config {
	repo := flag.String("path", "", "Path to files")

	mqttHost := flag.String("host", "localhost:1883", "MQTT broker host")
	mqttUser := flag.String("user", "", "MQTT broker host")
	mqttPassword := flag.String("password", "", "MQTT broker host")
	mqttBaseTopic := flag.String("topic", "glitter", "MQTT base topic")

	httpListen := flag.String("listen", ":8023", "HTTP Listen Port")

	gpioPin := flag.Int("gpio", 479, "GPIO pin")

	flag.Parse()

	mqttConfig := &MqttConfig{
		Host:     *mqttHost,
		User:     *mqttUser,
		Password: *mqttPassword,

		BaseTopic: *mqttBaseTopic,
	}

	httpConfig := &HttpConfig{
		Listen: *httpListen,
	}

	config := &Config{
		Mqtt:     mqttConfig,
		Http:     httpConfig,
		RepoPath: *repo,
		GpioPin:  *gpioPin,
	}

	return config
}
