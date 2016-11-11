package main
import(
    "fmt"
    "os"
)

func main(){
    dir, err := os.Open(".")
    if err != nil{panic("Cannot read dir");}
    defer dir.Close()
    // Readdir takes number of file. -1 = recursive
    fileInfo, err := dir.Readdir(5)
    if err != nil{panic("Cannot read dir 22");}
    for _, fi := range fileInfo{
        fmt.Println(fi.Name())
    }
}
