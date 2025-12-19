package main

import (
	"time"
)

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(2 * time.Second)
	println(phrase)
	doneChan <- true
}
func greet(phrase string, doneChan chan bool) {
	println(phrase)
	doneChan <- true
}
func main() {
	doneChan := make(chan bool)
	go greet("Nice to meet you!", doneChan)
	go greet("How are you?!", doneChan)
	go slowGreet("How...are...you?", doneChan)
	go greet("I Hope you´re liking studing Go!", doneChan)

	for i := 0; i < 4; i++ {
		<-doneChan
	}
	// fmt.Print((<-doneChan))
}
