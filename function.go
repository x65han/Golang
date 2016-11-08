package main
import "fmt"

func main(){
		list := []int{1,2,3,4,5,6,7,8,8}
		fmt.Println(list);
		fmt.Println(average(list));
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
