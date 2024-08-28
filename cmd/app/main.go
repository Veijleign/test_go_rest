package main

import (
  "encoding/json"
  "net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {

	http.HandleFunc("/", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{Message: "Hello World"}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}

}
