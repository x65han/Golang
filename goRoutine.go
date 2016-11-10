package main
import (
	"fmt"
	"time"
	"math/rand"
)
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Thread: ", n, " : State: ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
//Concurrency
func main() {
	defer cleanup(" :<<-|*|->>: ");
	i := 0
	for i < 10{
		go f(i);i++;
	} 
	var input string
	fmt.Scanln(&input)  
}
func cleanup(header string){
	fmt.Println(header);
}