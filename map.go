package main
import "fmt"
func main(){
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println("-=-=-=-=-=-=-");
	// Check if element exist
	if name,ok:=elements["Li"];ok{
		fmt.Println(name);
	}else{
		fmt.Println(ok)
	}
	// Delete element
	fmt.Println("-=-=-deleting Li=-=-=-");
	delete(elements, "Li");
	// Check if element exist
	if name,ok:=elements["Li"];ok{
		fmt.Println(name);
	}else{
		fmt.Println(ok)
	}
}