package main

import (
	"awesomeProject/accessModifiers"
	_ "awesomeProject/accessModifiers"
	"fmt"
)

func main() {
	// Hello world
	fmt.Println("Hello world")
	//public variable
	fmt.Println(accessModifiers.Version)
	//public getter func
	fmt.Println(accessModifiers.GetName())
}
