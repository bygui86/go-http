package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("*** using wrapper")
	wrapper()
	fmt.Println("")
	fmt.Println("*** using low level")
	lowLevel()
}

func wrapper() {
	// request
	reqBody := strings.NewReader(`
		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}
	`)
	request, _ := http.NewRequest(
		"POST",
		"http://dummy.restapiexample.com/api/v1/create",
		reqBody,
	)
	request.Header.Add("Content-Type", "application/json; charset=UTF-8")

	// send HTTP request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// response body
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("status: %d\n", response.StatusCode)
	fmt.Printf("body: %s\n", data)
}

func lowLevel() {
	// request
	postUrl, _ := url.Parse("http://dummy.restapiexample.com/api/v1/create")
	requestBody := ioutil.NopCloser(strings.NewReader(`
		{
			"name":"test",
			"salary":"123",
			"age":"23"
		}
	`))
	request := &http.Request{
		Method: "POST",
		URL:    postUrl,
		Header: map[string][]string{
			"Content-Type": {"application/json; charset=UTF-8"},
		},
		Body: requestBody,
	}

	// send HTTP request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// response body
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("status: %d\n", response.StatusCode)
	fmt.Printf("body: %s\n", data)
}
