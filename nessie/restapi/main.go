package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Read the index.html
	html, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println("Error reading index.html:", err)
		return
	}

	// Serve the OpenAPI spec
	http.HandleFunc("/nessie-openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, "nessie-openapi.yaml")
	})

	// Serve the index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	})

	fmt.Println("Serving nessie-openapi.yaml on http://localhost:19121/")

	// Start the server
	http.ListenAndServe(":19121", nil)
}
