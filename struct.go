package main
import ("fmt";"math")

type Circle struct{
	x,y,r float64
}
func main() {
	var c Circle;
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	// c := Circle{0, 0, 5}
	c.x = 1
	c.y = 2
	c.r = 3
	fmt.Println(circleArea(c))
}
func circleArea(c Circle) float64{
	return round(math.Pi * c.r * c.r)
}
func round(input float64) float64{
	return float64(int(input * 100)) / 100.0
}