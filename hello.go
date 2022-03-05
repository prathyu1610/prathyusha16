package main

import "fmt"

type restuarent struct {
	name   string
	cusine string
}

func main() {

	var restuarentslice []restuarent
	num := restuarent{"prathyusha", "indian"}
	num1 := restuarent{"shalu", "italian"}
	num2 := restuarent{name: "appu"}
	num3 := restuarent{"varsha", "chinese"}

	restuarentslice = append(restuarentslice, num, num1, num2, num3)
	fmt.Println(restuarentslice)

	var intslice []int
	intslice = append(intslice, 1, 2, 3)
	fmt.Println(intslice)
	intslice = append(intslice, 4, 5, 6)
	fmt.Println(intslice)
}
