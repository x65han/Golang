package main
import "fmt"
func main(){
	fmt.Print("Enter a float: ")
	var input float64
	fmt.Scanf("%f", &input);
	var output = input * 2;
	fmt.Println(output);

	fmt.Println("What is your name?");
	var name string;
	fmt.Scanf("%s", &name);
	fmt.Println("Your name is: " + name);
}