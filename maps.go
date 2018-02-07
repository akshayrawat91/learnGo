package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	// OR
	//lookup := map[string]int{	// composite literal
	//	"goku": 9001,
	//	"gohan": 4000,
	//}

	power, exists := lookup["vegeta"] // check if the key "vegeta" exists
	fmt.Println(power, exists)

	total := len(lookup) // number of keys in the map
	print(total)
	delete(lookup, "goku") // delete the value based on the key

	// initialize struct Saiyan
	//goku := &Saiyan{
	//	Name:"goku",
	//	Friends: make(map[string]*Saiyan),
	//}
	//
	//goku.Friends["krillin"] = ... // TODO

	// iteration on maps is in random order
	//for key, value := range lookup{
	//
	//}
}

//type Saiyan struct {
//	Name string
//	Friends map[string]*Saiyan
//}
