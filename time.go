package main
import(
	"fmt"
	"time"
)
func main(){
	t := time.Now()
	fmt.Println(t.Format("20060102150405"))
	fmt.Println(time.Now().Format(time.RFC850))
}
