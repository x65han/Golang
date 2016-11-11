package main

import(
    "fmt"
    "container/list"
    "reflect"
)

func main(){
    var x list.List
    x.PushBack(1)
    x.PushBack("Johnson")
    x.PushBack(3)

    for i := x.Front(); i != nil; i = i.Next(){
        fmt.Println(reflect.TypeOf(i.Value))
        fmt.Println(i.Value)
    }
}
