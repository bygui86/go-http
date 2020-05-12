package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// make a sample HTTP GET request
	response, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal(err)
	}

	// response headers
	fmt.Println(response.Header)
	contentType := response.Header.Get("Content-Type")
	fmt.Println("Content-Type", contentType)

	// response body
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("%s\n", data)
}
