package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	// client
	customClient := &http.Client{
		Timeout: 100 * time.Microsecond,
	}

	// send GET request
	response, err := customClient.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		// get `url.Error` struct pointer from `err` interface
		urlErr := err.(*url.Error)

		// check if error occurred due to timeout
		if urlErr.Timeout() {
			fmt.Println("Error occurred due to a timeout.")
		}

		// log error and exit
		log.Fatal("Error:", err)
	} else {
		fmt.Println("Success: status-code", response.StatusCode)
	}
}
