package main

import (
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"net/http"

	"fmt"
	"log"
)

type ApiError struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message error  `json:"message"`
}

func apiRegisterRoutes(router *httprouter.Router) *httprouter.Router {
	// Shader :: Meta
	router.GET("/api/v1/shaders", apiEndpoint(apiShaderList))
	router.POST("/api/v1/shaders", apiEndpoint(apiShaderCreate))
	router.GET("/api/v1/shaders/:id", apiEndpoint(apiShaderGet))
	router.PUT("/api/v1/shaders/:id", apiEndpoint(apiShaderUpdate))
	router.DELETE("/api/v1/shaders/:id", apiEndpoint(apiShaderDelete))

	// Shader :: Program
	router.GET("/api/v1/shaders/:id/program", apiEndpoint(apiShaderProgramGet))
	router.POST("/api/v1/shaders/:id/program", apiEndpoint(apiShaderProgramUpdate))
	router.PUT("/api/v1/shaders/:id/program", apiEndpoint(apiShaderProgramUpdate))
	router.DELETE("/api/v1/shaders/:id/program", apiEndpoint(apiShaderProgramDelete))

	return router
}

// Decorate API endpoints with logging and stuff
func apiEndpoint(handle httprouter.Handle) httprouter.Handle {
	return func(
		res http.ResponseWriter,
		req *http.Request,
		params httprouter.Params) {

		log.Println(req.Method, req.URL)

		handle(res, req, params)
	}
}

func apiWriteResponseJson(res http.ResponseWriter, payload interface{}) {
	res.Header().Add("Content-Type", "application/json")
	encoded, err := json.Marshal(payload)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("json_encoding_error", err))
		return
	}

	res.Write(encoded)
}

func apiEncodeError(errorType string, err error) []byte {
	result, _ := json.Marshal(ApiError{
		Status:  500,
		Error:   errorType,
		Message: err,
	})
	return result
}

// Shader Meta / File API
func apiShaderList(
	res http.ResponseWriter,
	req *http.Request,
	_ httprouter.Params) {

	fmt.Fprintf(res, "apiShaderList")
}

func apiShaderCreate(
	res http.ResponseWriter,
	req *http.Request,
	_ httprouter.Params) {

	fmt.Fprintf(res, "apiShaderCreate")
}

func apiShaderGet(
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	fmt.Fprintf(res, "apiShaderGet")
}

func apiShaderUpdate(
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	fmt.Fprintf(res, "apiShaderUpdate")
}

func apiShaderDelete(
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	fmt.Fprintf(res, "apiShaderDelete")
}

// Shader Program API

func apiShaderProgramGet(
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	fmt.Fprintf(res, "apiShaderProgramGet")
}

func apiShaderProgramUpdate(
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	fmt.Fprintf(res, "apiShaderProgramUpdate")
}

func apiShaderProgramDelete(
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	fmt.Fprintf(res, "apiShaderProgramDelete")
}
