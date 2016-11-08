package main
import "fmt"

func main(){
	var x = 10
	determine(x)
	x--
	determine(x)
	x--
	x--
	determine(x)
}
func determine(x int){
	if x % 2 == 0{
		fmt.Println(x, " is divisible by 2");
	}else if x % 3 == 0{
		fmt.Println(x, " is divisible by 3");
	}else{
		fmt.Println("Not 2 or 3");
	}
}