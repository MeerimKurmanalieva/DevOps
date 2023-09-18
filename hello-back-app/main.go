package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a HelloRequest struct
	var request HelloRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Create a response message
	message := fmt.Sprintf("Hello, %s!", request.Name)

	// Create a JSON response
	response := HelloResponse{
		Message: message,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
