package main

import (
	"fmt"
	"net/http"
	"os"
)

type HelloRequest struct {
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/api/hello" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	name := r.URL.Query().Get("name") // Extract the name from query parameter

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Name parameter is missing")
		return
	}

	response := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, response)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cli <host:port>")
		return
	}

	host := os.Args[1]

	http.HandleFunc("/api/hello", helloHandler)

	// Check if there is a command-line argument for name
	if len(os.Args) >= 3 {
		name := os.Args[2]
		response := fmt.Sprintf("Hello, %s!", name)
		fmt.Println(response)
	}

	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
