package main

import (
	"awesomeProject/accessModifiers"
	"awesomeProject/enum"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
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

}
