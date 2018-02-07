package main

import (
	"fmt"
)

type Saiyan struct {
	name  string
	power int
}

//type Saiyan2 struct {
//	name string
//	power int
//	father *Saiyan2
//}

func (s *Saiyan) Super() { // *Saiyan is receiver of Super method
	s.power += 10001
}

type Person struct {
	name string
}

func (s *Person) Interface() {
	fmt.Printf("name is %s", s.name)
}

type Saiyaman struct {
	*Person // composition
	power   int
}

func (s *Saiyaman) Interface() { // overriding is not present in Go so using composition take place of previous func
	fmt.Printf("\n name is %s", s.name)
}

func newSuper(name string, power int) Saiyan {
	return Saiyan{
		name:  name,
		power: power,
	}
}

func main() {
	//x := 100
	/*if len(os.Args) != 2{
		os.Exit(1)
	}*/

	goku := &Saiyan{
		name:  "goku",
		power: 9000,
	}
	// OR
	gohan := new(Saiyan)
	gohan.name = "gohan"
	gohan.power = 4000

	//newSuper(gohan.name, gohan.power)
	goku.Super()
	//Super(goku)

	vegeta := &Saiyaman{
		Person: &Person{name: "vegeta"},
		power:  8000,
	}

	vegeta.Interface()

	fmt.Printf("\n its over %d, %d", goku.power, gohan.power)
	fmt.Printf("\n i am %s", vegeta.name)

	//var ar [4]int		// declaring array
	ar1 := [4]int{0, 1, 2, 3} // another way to declare and initialize an array
	//println("\n",ar1[1])

	//slic := []int{20,21,23}		// declare and initialize a slice
	//slic := make([]int,10)		// another way to declare a slice (length and capacity of 10)
	slic := make([]int, 0, 10) // length of 0 and capacity of 10
	//c := cap(slic)
	//println(c)
	slic = append(slic, 5)

	println("\n", slic[0])
	println(len(slic))

}

//func Super(s *Saiyan)  {
//	s.power += 10000
//}
