package main
import ("fmt" ; "sort")
type Person struct {
    Name string
    Age int
}
// ByName
type ByName []Person
func (this ByName) Len() int {
    return len(this)
}
func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}
// ByAge
type ByAge []Person
func (this ByAge) Len() int {
    return len(this)
}
func (this ByAge) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}
func main() {
    kids := []Person{
                {"Zack",9},
                {"Jack",19},
                {"Aaron",10},
    }
    fmt.Println("Before sort:\n", kids)
    sort.Sort(ByName(kids))
    fmt.Println("By name:\n", kids)
    sort.Sort(ByAge(kids))
    fmt.Println("By age:\n", kids)
}
