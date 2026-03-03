package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Message: "Hello World",
		Time:    time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server chạy tại port 8080...")
	http.ListenAndServe(":8080", nil)
}
