package main

import (
	"fmt"
	"net/http"
)

var version = "unknown"

func main() {
	fmt.Println("Glitter\t\t\t\t\t\t", version)

	http.ListenAndServe("localhost:8023", nil)
}
