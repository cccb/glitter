package main

import (
	"github.com/julienschmidt/httprouter"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"
	"log"
	"strconv"
)

type ApiError struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message error  `json:"message"`
}

type ApiContext struct {
	ShaderRepo *ShaderRepository
}

type ApiHandle func(
	context *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params)

func apiRegisterRoutes(
	context *ApiContext,
	router *httprouter.Router,
) *httprouter.Router {

	// Shader :: Meta
	router.GET("/api/v1/shaders", apiEndpoint(context, apiShaderList))
	router.POST("/api/v1/shaders", apiEndpoint(context, apiShaderCreate))
	router.GET("/api/v1/shaders/:id", apiEndpoint(context, apiShaderGet))
	router.PUT("/api/v1/shaders/:id", apiEndpoint(context, apiShaderUpdate))
	router.DELETE("/api/v1/shaders/:id", apiEndpoint(context, apiShaderDelete))

	// Shader :: Program
	router.GET("/api/v1/shaders/:id/program", apiEndpoint(context, apiShaderProgramGet))
	router.POST("/api/v1/shaders/:id/program", apiEndpoint(context, apiShaderProgramUpdate))
	router.PUT("/api/v1/shaders/:id/program", apiEndpoint(context, apiShaderProgramUpdate))

	return router
}

// Decorate API endpoints with logging and stuff
func apiEndpoint(ctx *ApiContext, handle ApiHandle) httprouter.Handle {
	return func(
		res http.ResponseWriter,
		req *http.Request,
		params httprouter.Params) {

		// Set Headers
		res.Header().Add("Content-Type", "application/json")

		// Log Request
		log.Println(req.Method, req.URL)

		handle(ctx, res, req, params)
	}
}

func apiWriteResponseJson(res http.ResponseWriter, payload interface{}) {
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
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	_ httprouter.Params) {

	shaders, err := ctx.ShaderRepo.List()
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("shader_list_error", err))
		return
	}

	apiWriteResponseJson(res, shaders)
}

func apiShaderCreate(
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	_ httprouter.Params) {

	var shader *Shader

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	if err := decoder.Decode(&shader); err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("parse_error", err))
		return
	}

	// Create Shader
	id, err := ctx.ShaderRepo.Create(shader)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("parse_error", err))
		return
	}

	shader.Id = id

	apiWriteResponseJson(res, shader)
}

func apiShaderGet(
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {
	shaderId, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("param_error", err))
		return
	}

	shader, err := ctx.ShaderRepo.Find(shaderId)
	if err != nil {
		// This most likely happened because of an invalid id
		// so, let's respond with a 404
		res.WriteHeader(404)
		res.Write(apiEncodeError(
			"shader_not_found", fmt.Errorf("Shader could not be found.")))
		return
	}

	// Strip token from shader
	shader.Token = ""

	apiWriteResponseJson(res, shader)
}

func apiShaderUpdate(
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {
	shaderId, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("param_error", err))
		return
	}

	var shader *Shader

	// Decode shader
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	if err := decoder.Decode(&shader); err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("parse_error", err))
		return
	}

	err = ctx.ShaderRepo.Update(shaderId, shader)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("update_error", err))
	}

	apiWriteResponseJson(res, shader)
}

func apiShaderDelete(
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {
	shaderId, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("param_error", err))
		return
	}

	shader, err := ctx.ShaderRepo.Find(shaderId)
	if err != nil {
		// This most likely happened because of an invalid id
		// so, let's respond with a 404
		res.WriteHeader(404)
		res.Write(apiEncodeError(
			"shader_not_found", fmt.Errorf("Shader could not be found.")))
		return
	}

	err = shader.Destroy()
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("delete_error", err))
	}

	apiWriteResponseJson(res, "SUCCESS")
}

//
// Shader Program API
//

/*
 GET ShaderProgram
*/
func apiShaderProgramGet(
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {
	shaderId, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("param_error", err))
		return
	}

	shader, err := ctx.ShaderRepo.Find(shaderId)
	if err != nil {
		// This most likely happened because of an invalid id
		// so, let's respond with a 404
		res.WriteHeader(404)
		res.Write(apiEncodeError(
			"shader_not_found", fmt.Errorf("Shader could not be found.")))
		return
	}

	program := shader.Program()
	if err != nil {
		// Keep things simple for now:
		// Always return an empty program
		program = []byte{}
	}

	res.Write(program)
}

/*
 POST / PUT ShaderProgram
*/
func apiShaderProgramUpdate(
	ctx *ApiContext,
	res http.ResponseWriter,
	req *http.Request,
	params httprouter.Params) {

	shaderId, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("param_error", err))
		return
	}

	program, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("parse_error", err))
		return
	}

	shader, err := ctx.ShaderRepo.Find(shaderId)
	if err != nil {
		// This most likely happened because of an invalid id
		// so, let's respond with a 404
		res.WriteHeader(404)
		res.Write(apiEncodeError(
			"shader_not_found", fmt.Errorf("Shader could not be found.")))
		return
	}

	err = shader.UpdateProgram(program)
	if err != nil {
		res.WriteHeader(500)
		res.Write(apiEncodeError("program_update_error", err))
		return
	}

	res.Write(program)
}
