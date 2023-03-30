package main

import (
	"awesomeProject/accessModifiers"
	enum "awesomeProject/enum"
	"fmt"
	"reflect"
	"strconv"
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
	myInt := 64
	myFloat := float64(myInt)
	fmt.Println(reflect.TypeOf(myInt), "to", reflect.TypeOf(myFloat))
}
