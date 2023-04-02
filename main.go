package main

import (
	"awesomeProject/accessModifiers"
	"awesomeProject/enum"
	"fmt"
	"reflect"
	"strconv"
	"time"
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
}
