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
	if len(os.Args) < 4 {
		fmt.Println("Usage: cli <hostname:port> <action> <name>")
		os.Exit(1)
	}

	backendURL := os.Args[1]
	action := os.Args[2]
	name := os.Args[3]

	var endpoint string

	switch action {
	case "create":
		endpoint = "/api/ec2/create"
	case "delete":
		endpoint = "/api/ec2/terminate"
	default:
		fmt.Println("Invalid action. Use 'create' or 'delete'.")
		os.Exit(1)
	}

	// Create a request payload
	payload := map[string]string{"name": name}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON payload:", err)
		os.Exit(1)
	}

	// Send a POST request to the appropriate endpoint
	url := fmt.Sprintf("http://%s%s", backendURL, endpoint)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
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
