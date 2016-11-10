package main
import "fmt"

func setToOne(xptr *int){
	*xptr = 1;
}
func main(){
	xptr := new(int)
	setToOne(xptr);
	fmt.Println(*xptr);
}