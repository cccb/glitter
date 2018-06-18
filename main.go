package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var version = "unknown"

func main() {
	fmt.Println("Glitter\t\t\t\t\t\t", version)

	router := httprouter.New()
	apiRegisterRoutes(router)

	// Welcome Page
	router.GET("/", func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		fmt.Fprintf(res, "WELCOME")
	})

	http.ListenAndServe("localhost:8023", router)
}
