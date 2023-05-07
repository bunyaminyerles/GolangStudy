package main

import (
	"fmt"
	"net/http"
)

func handleRequests(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handleRequests)
	http.ListenAndServe(":9000", nil)
	fmt.Println("Server is running on port 9000")
}
