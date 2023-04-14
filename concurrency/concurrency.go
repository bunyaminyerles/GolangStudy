package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(8)
	go println("World")
	println("Hello")
	go println("a")
	go println("b")
	go println("c")
	time.Sleep(100 * time.Millisecond)
}
