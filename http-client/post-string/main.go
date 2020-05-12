package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// request body (payload)
	requestBody := strings.NewReader(`
		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}
	`)

	// post some data
	response, err := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"application/json; charset=UTF-8",
		requestBody,
	)
	if err != nil {
		log.Fatal(err)
	}

	// response headers
	requestContentType := response.Request.Header.Get("Content-Type")
	fmt.Println("Request content-type:", requestContentType)

	// response data
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("%s\n", data)
}
