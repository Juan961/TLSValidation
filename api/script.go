package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func script() {
	const baseURL string = "https://api.ssllabs.com/api/v2/analyze?host=www.ssllabs.com"

	resp, err := http.Get(baseURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
