package main
import "fmt"
//e.g. close a file when function terminates
//Useful when there is mutiple return statements
//Like Java finally clause, run even there is exception
func main(){
	defer second();
	first();
}
func first(){
	fmt.Println("first");
}
func second(){
	fmt.Println("second");
}