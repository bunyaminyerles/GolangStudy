package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	Name     string
	Lastname string
	Age      int
}

func (h Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Name = "John"
	h.Lastname = "Doe"
	h.Age = 30

	// form parsing
	r.ParseForm()

	// form data
	fmt.Println(r.Form)

	// url path
	fmt.Println("Path: ", r.URL.Path)

	fmt.Fprintf(w, "<table><tr><td>Name</td><td>Lastname</td><td>Age</td></tr><tr><td>%s</td><td>%s</td><td>%d</td></tr></table>", h.Name, h.Lastname, h.Age)

}

func main() {
	var h Human
	err := http.ListenAndServe("localhost:9000", h)
	checkError(err)
	fmt.Println("Server is running on port 9000")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
