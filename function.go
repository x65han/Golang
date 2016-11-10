package main
import "fmt"

func main(){
		list := []int{1,2,3,4,5,6,7,8,8}
		fmt.Println(list);
		fmt.Println("Average: ", average(list));
		fmt.Println(rint());
		// internal Function
		add := func(x,y int) int{
				return x + y
		}
		fmt.Println(add(7,6));
}
func average(input []int) float64{
		total := 0.0
		for _,x := range input{
				total += float64(x)
		}
		return round(total/float64(len(input)) )
}
func round(input float64) float64{
		return float64(int(input * 100)) / 100.0
}
//function with named return type
func rint() (x,y,z int){
		x = 1
		y = 2
		z = 3
		return
}
