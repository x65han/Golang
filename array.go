package main
import "fmt"
func main(){
	// Create Array
	arr := []int{1997,1998,1999,2000,2001}
	fmt.Println(arr);
	// Create Slice Array
	sliceArr := make([]int,3)
	fmt.Println(sliceArr);
	// Append Array
	appendArr := append(sliceArr, 2016, 2017);
	fmt.Println(appendArr);
	// Copy Array
	appendArr[0] = 2013
	appendArr[1] = 2014
	appendArr[2] = 2015
	copy(arr, appendArr);
	fmt.Println(arr);
}
