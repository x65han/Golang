package main
import "fmt"
func main(){
	var x string = "Hello"
	var y string
	y = x + " Earth"
	fmt.Println(x);
	fmt.Println(y);
	//Easy Declaration
	name := "Not Born Yet"
	var age = -1;
	fmt.Println("My dogs name is: " + name);
	fmt.Println("Age is: ", age);
	//Mass Declaration 
	author, school, company := "Johnson Han", "University of Waterloo", "BlackBerry"
	fmt.Println("The author " + author + " goes to " + school + " and is currently working for " + company);
}