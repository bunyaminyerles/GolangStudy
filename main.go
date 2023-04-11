package main

import (
	"awesomeProject/accessModifiers"
	"awesomeProject/enum"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	// Hello world
	fmt.Println("Hello world")

	// public variable
	fmt.Println(accessModifiers.Version)

	// public getter func
	fmt.Println(accessModifiers.GetName())

	// Enum example
	enum.PrintBrand(enum.BMW)
	enum.PrintBrand(enum.MERCEDES)

	// type convert
	// string to int
	myString1 := "123"
	number, _ := strconv.Atoi(myString1)
	fmt.Println(reflect.TypeOf(myString1), "to", reflect.TypeOf(number))
	// string to bool
	myString2 := "false"
	myBoolean, _ := strconv.ParseBool(myString2)
	fmt.Println(reflect.TypeOf(myString2), "to", reflect.TypeOf(myBoolean))

	// type casting
	myInt1 := 64
	myFloat1 := float64(myInt1)
	fmt.Println(reflect.TypeOf(myInt1), "to", reflect.TypeOf(myFloat1))

	// formatter
	// https://pkg.go.dev/fmt
	myString3 := "alfa"
	fmt.Printf("Data types : %T \n", myString3)
	fmt.Printf("Value of myString3 variable : %v \n", myString3)

	/*
		// input operations from console
		reader := bufio.NewReader(os.Stdin)
		myString4, _ := reader.ReadString('\n')
		fmt.Println("input data : ", myString4)
		// input number value
		myString5, _ := reader.ReadString('\n')
		myFloat2, err := strconv.ParseFloat(strings.TrimSpace(myString5), 64)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Value of number : ", myFloat2)
		}
	*/

	// Date and hour operations
	specificDate := time.Date(2023, time.April, 3, 0, 40, 0, 0, time.UTC)
	fmt.Printf("Specific date : %s\n", specificDate)
	fmt.Println("The month is ", specificDate.Month())
	fmt.Println("The day is ", specificDate.Day())
	fmt.Println("The weekday is ", specificDate.Weekday())
	nextYear := specificDate.AddDate(1, 0, 0)
	fmt.Printf("The next year same time is : %s\n", nextYear)
	longFormat := "Monday, January 2, 2006"
	fmt.Println("Long format : ", nextYear.Format(longFormat))
	shortformat := "2/1/2006"
	fmt.Println("Short format :", nextYear.Format(shortformat))
	fmt.Println("Timestamp : ", nextYear.Unix())
	now := time.Now()
	fmt.Printf("The time now is : %s\n", now)

	// if condition block
	a := 20
	b := 10

	if a < b {
		fmt.Println("a smaller than b")
	} else if a > b {
		fmt.Println("a bigger than b")
	} else {
		fmt.Println("a equals b")
	}
	// variable definition at if scope
	if x, y := 10, 20; x < y {
		fmt.Println("Variables of x and y can only be called in the if scope.")
	}

	// switch condition
	city := "Ankara"
	switch city {
	case "Ankara":
		fmt.Println("This city is Ankara")
	case "Istanbul":
		fmt.Println("This city is Istanbul")
	default:
		fmt.Println("This city is not found")
	}
	grade := 75
	switch {
	case grade >= 0 && grade <= 59:
		fmt.Println("Your alphanumeric grade is F")
	case grade > 59 && grade <= 69:
		fmt.Println("Your alphanumeric grade is D")
	case grade > 69 && grade <= 79:
		fmt.Println("Your alphanumeric grade is C")
	case grade > 79 && grade <= 89:
		fmt.Println("Your alphanumeric grade is B")
	case grade > 89 && grade <= 100:
		fmt.Println("Your alphanumeric grade is A")
	default:
		fmt.Println("Your grade is invalid")
	}

	// break, continue, goto keywords
	// break
	fmt.Print("i values :")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Print(i, "\t")
		i++
	}
	fmt.Print("\n")
	// continue
	fmt.Print("i values :")
	for i := 0; i < 7; i++ {
		if i%3 == 0 {
			continue
		}
		if i == 5 {
			break
		}
		fmt.Print(i, "\t")
	}
	fmt.Print("\n")
	// goto
	fmt.Print("i values :")
	for i := 0; i < 7; i++ {
		if i%3 == 0 {
			continue
		}
		goto PRINT
		if i == 5 {
			break
		}
	PRINT:
		fmt.Print(i, "\t")
	}
	fmt.Print("\n")

	// loop
	// for
	fmt.Println("Loop \nfor")
	fmt.Print("Values : ")
	for i := 0; i <= 5; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Print("\n")
	// while
	fmt.Println("while")
	fmt.Print("Values : ")
	j := 1
	for j < 10 {
		fmt.Print(j, "\t")
		j += j
	}
	fmt.Print("\n")
	// range array example
	fmt.Println("Print the exponents of 2 with 'range'.")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// range map example
	capitals := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Rome"}
	for country := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capitals[country])
	}
	// range string example
	myString6 := "golang"
	for i := range myString6 {
		fmt.Println(i, myString6[i])
	}

	// error handling
	result, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// create error object
	myError := errors.New("This is my error")
	fmt.Println(myError.Error())

	// array
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "yellow")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	// array with make
	myArray := make([]int, 5, 5)
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5
	fmt.Println(myArray)
	fmt.Println("append before len : ", len(myArray))
	fmt.Println("append before cap : ", cap(myArray))
	myArray = append(myArray, 6)
	fmt.Println(myArray)
	fmt.Println("append after len : ", len(myArray))
	fmt.Println("append after cap : ", cap(myArray))
	// array with initialization
	myArray2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray2)
	// array with initialization and ...
	myArray3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(myArray3)
	// array with initialization and index
	myArray4 := [...]int{1: 1, 3: 2, 5: 3}
	fmt.Println(myArray4)
	// slice
	mySlice := myArray[1:4]
	fmt.Println(mySlice)
	mySlice2 := myArray[1:]
	fmt.Println(mySlice2)
	mySlice3 := myArray[:4]
	fmt.Println(mySlice3)

	// map
	myMap := make(map[string]string)
	myMap["name"] = "golang"
	myMap["version"] = "1.14.2"
	fmt.Println(myMap)
	fmt.Println(myMap["name"])
	fmt.Println(myMap["version"])
	// map with initialization
	myMap2 := map[string]string{"name": "golang", "version": "1.14.2"}
	fmt.Println(myMap2)
	// delete map
	delete(myMap2, "name")
	fmt.Println(myMap2)
	// map with struct
	type Person struct {
		name    string
		surname string
		age     int
	}
	myMap3 := make(map[string]Person)
	myMap3["person1"] = Person{"golang", "gopher", 1}
	myMap3["person2"] = Person{"python", "snake", 2}
	fmt.Println(myMap3)
	fmt.Println(myMap3["person1"])
	fmt.Println(myMap3["person2"])
	// map traverse
	for key, value := range myMap3 {
		fmt.Println(key, value)
	}
	// map traverse with key
	list := make([]string, 0, len(myMap3))
	for key := range myMap3 {
		list = append(list, key)
	}
	fmt.Println(list)
	for _, v := range list {
		fmt.Println(myMap3[v])
	}

	// strings package
	fmt.Print("\nStrings package \n")
	fmt.Println(strings.Contains("golang", "go"))
	fmt.Println(strings.Count("golang", "o"))
	fmt.Println(strings.HasPrefix("golang", "go"))
	fmt.Println(strings.HasSuffix("golang", "lang"))
	fmt.Println(strings.Index("golang", "l"))
	fmt.Println(strings.Join([]string{"golang", "python", "java"}, "-"))
	fmt.Println(strings.Repeat("golang\t", 3))
	fmt.Println(strings.Replace("golang golang", "go", "python", 2))
	fmt.Println(strings.Split("golang-python-java", "-"))
	fmt.Println(strings.ToLower("GOLANG"))
	fmt.Println(strings.ToUpper("golang"))
	fmt.Println(strings.TrimSpace(" golang "))
	fmt.Println(strings.Trim("golang", "g"))
	fmt.Println(strings.TrimLeft("golang", "g"))
	fmt.Println(strings.TrimRight("golang", "g"))
	fmt.Println(strings.Fields("golang python java"))
	fmt.Println(strings.FieldsFunc("golang,python,java", func(r rune) bool {
		return r == ','
	}))
	fmt.Println(strings.SplitAfter("golang-python-java", "-"))
	fmt.Println(strings.SplitAfterN("golang-python-java", "-", 2))
	fmt.Println(strings.SplitN("golang-python-java", "-", 2))
	fmt.Println(strings.Title("golang python java"))
	fmt.Println(strings.ToTitle("golang python java"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.TrimPrefix("golang", "go"))
	fmt.Println(strings.TrimSuffix("golang", "lang"))
	fmt.Println(strings.Compare("golang", "golang"))
	fmt.Println(strings.EqualFold("golang", "GOLANG"))
	fmt.Println(strings.NewReplacer("golang", "python", "python", "java").Replace("golang python"))
	fmt.Println(strings.ReplaceAll("golang python", "golang", "python"))

}
