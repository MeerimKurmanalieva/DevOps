package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: cli <hostname:port> <name>")
		os.Exit(1)
	}

	backendURL := os.Args[1]
	name := os.Args[2]

	// Create a request payload
	payload := map[string]string{"name": name}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON payload:", err)
		os.Exit(1)
	}

	// Send a GET request to the backend
	url := fmt.Sprintf("http://%s/api/hello", backendURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
