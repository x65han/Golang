package main
import "fmt"
func main(){
	var x = 0;
	swi(x)
	x++
	swi(x)
	x++
	swi(x)
	x++
}
func swi(x int){
	// No need to add break in between
	switch x{
		case 0: fmt.Println("0 is printed");
		case 1: fmt.Println("1 is printed");
		default: fmt.Println("Switch Default");
	}
}