package main

import (
	"sync"
)

var c = make(chan int)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go xFunc(c, wg)
	go printer(c, wg)
	defer wg.Wait()
}

func xFunc(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for l := 1; l < 20; l++ {
		c <- l
	}
	close(c)
}

func printer(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case message, ok := <-c:
			if !ok {
				return
			}
			println(message)
		}
	}
}
