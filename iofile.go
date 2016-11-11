package main
import(
    "fmt"
    "io/ioutil"
    "os"
)

func main(){
    // Read File
    bs, err := ioutil.ReadFile("input.txt");
    if err != nil{panic("Cannot read file");}
    str := string(bs)
    fmt.Println("\n" + str)
    // Write File
    file, err := os.Create("output.txt")
    if err != nil{panic("Cannot write file");}
    defer file.Close()
    file.WriteString(">> Writing\n")
    file.WriteString(str)
}
