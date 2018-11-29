package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cameliot/alpaca"
	"github.com/cameliot/alpaca/meta"
	"github.com/julienschmidt/httprouter"
)

var version = "unknown"

func usage() {
	flag.PrintDefaults()
	os.Exit(-1)
}

func startMqtt(config *Config) {
	// Setup MQTT
	actions, dispatch := alpaca.DialMqtt(
		config.Mqtt.BrokerUri(),
		alpaca.Routes{
			"power": config.Mqtt.BaseTopic + "/power",
			"meta":  "v1/_meta",
		})

	powerActions := make(alpaca.Actions)
	metaActions := make(alpaca.Actions)

	// Handle powersupply actions

	// Hanlde meta actions for service discovery
	metaSvc := meta.NewMetaSvc(
		"treppe@mainhall",
		"glitter",
		version,
		"Glitter MQTT Treppe",
	)
	go metaSvc.Handle(metaActions, dispatch)

	for action := range actions {
		powerActions <- action
		metaActions <- action
	}

}

func main() {
	fmt.Println("Glitter\t\t\t\t\t\t", version)

	// Initialize configuration
	config := parseFlags()
	if config.RepoPath == "" {
		usage()
	}

	// Setup Repository
	shaderRepo, err := NewShaderRepository(config.RepoPath)
	if err != nil {
		log.Panic("Could not initialize / use shader repository:", err)
	}

	// Setup MQTT
	go startMqtt(config)

	// Setup HTTP API
	router := httprouter.New()
	apiRegisterRoutes(&ApiContext{
		ShaderRepo: shaderRepo,
	}, router)

	// Welcome Page
	router.GET("/",
		func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
			fmt.Fprintf(res, "WELCOME")
		})

	log.Println("Listening for HTTP connections on", config.Http.Listen)
	http.ListenAndServe(config.Http.Listen, router)
}
