package main
import "fmt"
func main(){
	// Create Array
	var arr [5]int
	arr[0] = 1997
	arr[1] = 1998
	arr[2] = 1999
	arr[3] = 2000
	arr[4] = 2001
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
	copy(sliceArr, appendArr);
	fmt.Println(sliceArr);
}
