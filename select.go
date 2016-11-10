package main
import ("fmt";"time")

/* 
	Go has a special statement called "select" which works like a "switch" but for channels:
*/  

func main() {
	// Asynchronous Channel(Buffered Channels)
	// c := make(chan int, 1) // with capacity
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "First"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "Second"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for { 
			select {
				case msg1 := <- c1:
					fmt.Println("Message 1: ", msg1)
				case msg2 := <- c2:
					fmt.Println("Message 2: ", msg2)
				//time.After = Interval with new channel
				case <- time.After(time.Second):
					fmt.Println("timeout")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}