package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get an image file
	response, err := http.Get("https://bit.ly/2IRnmVm")
	if err != nil {
		log.Fatal(err)
	}

	// read response data
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	// check Content-Type
	if response.Header.Get("Content-Type") == "image/jpeg" {
		wrErr := ioutil.WriteFile("keyboard.jpeg", data, 0777)
		if wrErr != nil {
			log.Fatal("Error saving image to disk")
		}
		fmt.Println("File saved!")
	} else {
		log.Fatal("Error: Response is not an image.")
	}
}
