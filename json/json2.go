package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func getPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func getPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func getPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func writeMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("******************************")
}

func WriteEndLine() {
	fmt.Println("------------------------------")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func SaveJsonToFile(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func main() {
	person := Person{
		ID:        1,
		FirstName: "John",
		LastName:  "Smith",
		UserName:  "jsmith",
		Gender:    "Male",
		Name: Name{
			Family:   "Smith",
			Personal: "John",
		},
		Email: []Email{
			Email{ID: 1, Kind: "home", Address: "john@home.com"},
			Email{ID: 2, Kind: "work", Address: "john@work.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Music"},
			Interest{ID: 2, Name: "Basketball"},
		},
	}

	writeMessage("Reading operation started")

	writeMessage("Reading person's name")
	WriteStarLine()
	res := getPerson(&person)
	writeMessage(res)
	WriteEndLine()
	writeMessage("\n")
	writeMessage("Reading person's email address")
	WriteStarLine()
	res = getPersonEmailAddress(&person, 1)
	writeMessage(res)
	WriteStarLine()

	writeMessage("\n")

	writeMessage("Reading person's email with index")
	WriteStarLine()
	resEmail := getPersonEmailAddress(&person, 1)
	writeMessage(resEmail)
	WriteStarLine()

	writeMessage("\n")

	writeMessage("Personal email object with index")
	WriteStarLine()
	resEmail2 := getPersonEmail(&person, 1)
	writeMessage(resEmail2.Address)
	WriteStarLine()

	writeMessage("\n")

	writeMessage("Writing JSON to file")
	WriteStarLine()
	SaveJsonToFile("person.json", person)
	WriteStarLine()

	writeMessage("\n")

	writeMessage("Reading JSON from file")
	WriteStarLine()
	file, err := os.Open("person.json")
	checkError(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	person2 := Person{}
	err = decoder.Decode(&person2)
	checkError(err)
	fmt.Println(person2)
	WriteStarLine()

	writeMessage("\n")

}
