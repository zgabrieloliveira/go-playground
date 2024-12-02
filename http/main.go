package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// dois endpoints simples, com uma mensagem de texto e um JSON, respectivamente
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/status", handleStatus)
	// pra todas as outras páginas, vai retornar 404 not found
	http.HandleFunc("/", http.NotFoundHandler().ServeHTTP)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "Server is running",
		"code":   "200",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
