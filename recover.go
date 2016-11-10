package main
import "fmt"

func main(){
	defer func(){
		fmt.Println("defer function");
		str := recover()
		fmt.Println(str)
	}()
	panic("paniccingg")
	fmt.Println("end of main");
}