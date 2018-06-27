package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var version = "unknown"

func usage() {
	flag.PrintDefaults()
	os.Exit(-1)
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
