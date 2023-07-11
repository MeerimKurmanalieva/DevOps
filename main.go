package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// add new struct type *HelloRequest*
type HelloRequest struct {
	Name string `json:"name"`
}

// Updated the *helloHandler*
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/api/hello" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var req HelloRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request payload")
		return
	}

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Name parameter is missing")
		return
	}

	response := fmt.Sprintf("Hello, %s!", req.Name)
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.ListenAndServe(":8081", nil)
}
