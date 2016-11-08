package main

import "fmt"

//this is a comment

func main(){
	fmt.Println("Running");
	defer func(){
		str := recover();
		fmt.Println(str);
	}()
	panic("PANIC")
}
