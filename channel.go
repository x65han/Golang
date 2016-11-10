package main
import (
	"fmt"
	"time"
)
/*
	c chan <- string
	Now c can receive message
*/
func pinger(c chan <- string) {
	for {
		c <- "ping"
	}
}
func ponger(c chan string) {
	for { 
		c <- "pong"
	}
}
func printer(c chan string) {
	for {
		fmt.Println(<- c)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}