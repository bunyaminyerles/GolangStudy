package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type API2 struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageConcat := "User id: " + id
	message := API2{messageConcat}
	output, err := json.Marshal(message)
	checkError3(err)
	fmt.Fprintf(w, string(output))
}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{id:[0-9]+}", Hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":9000", nil)

}

func checkError3(err error) {
	if err != nil {
		panic(err)
	}
}
