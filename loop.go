package main
import "fmt"
func main(){
	var i = 1
	//while loop
	for i < 11 {
		fmt.Print(i, " ")
		i++;
	}
	fmt.Print("\n")

	for x := 0;x < 10;x++{
		fmt.Print(x, " ");
	}
	fmt.Print("\n");
}