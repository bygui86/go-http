package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type SampleRequest struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

type SampleResponse struct {
	Status string             `json:"status"`
	Data   SampleResponseData `json:"data"`
}

type SampleResponseData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

func main() {
	// request body (payload)
	requestBody, _ := json.Marshal(&SampleRequest{
		Name:   "test",
		Salary: "123",
		Age:    "23",
	})

	// post some data
	response, reqErr := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"application/json; charset=UTF-8",
		bytes.NewBuffer(requestBody),
	)
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	// response headers
	requestContentType := response.Request.Header.Get("Content-Type")
	fmt.Println("Request content-type:", requestContentType)

	// response data
	var responseBody SampleResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&responseBody)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}
	fmt.Printf("%+v\n", responseBody)

	resCloseErr := response.Body.Close()
	if resCloseErr != nil {
		log.Fatal(resCloseErr)
	}
}
