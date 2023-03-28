package main

import (
	"awesomeProject/accessModifiers"
	enum "awesomeProject/enum"
	"fmt"
)

func main() {
	// Hello world
	fmt.Println("Hello world")
	//public variable
	fmt.Println(accessModifiers.Version)
	//public getter func
	fmt.Println(accessModifiers.GetName())
	//Enum example
	enum.PrintBrand(enum.BMW)
	enum.PrintBrand(enum.MERCEDES)
}
