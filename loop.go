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
		//for loop
		for x := 0;x < 10;x++{
			fmt.Print(x, " ");
		}
		fmt.Print("\n");
		//for each loop
		arr := []int{1997, 1998, 1999, 2000}
		for index, x := range arr{
				fmt.Println("Index: ", index, " Value: ", x);
		}
}
