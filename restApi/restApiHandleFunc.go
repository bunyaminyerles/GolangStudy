package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {

	apiRoot := "/api"

	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError2(err)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, FirstName: "John", LastName: "Doe", Age: 25},
			{ID: 2, FirstName: "Jane", LastName: "Doe", Age: 30},
		}
		message := users
		output, err := json.Marshal(message)
		checkError2(err)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		message := User{ID: 1, FirstName: "John", LastName: "Doe", Age: 25}
		output, err := json.Marshal(message)
		checkError2(err)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(output))
	})

	http.ListenAndServe(":9000", nil)

}

func checkError2(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
