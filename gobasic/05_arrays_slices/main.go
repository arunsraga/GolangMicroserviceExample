package main


import "fmt"

func main(){
	// Arrays
	// slice is an array without any type

	var fruitArr [2]string

	//assign value
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	fmt.Println(fruitArr)

	// declare and assign array
	fruits := [2]string{"orange", "pineapple"}
	fmt.Println(fruits)


	// Slices
	fruitSlices := []string{"apple", "orange", "grape"}
	fmt.Println(fruitSlices)
	fmt.Println("length of fruitSlices", len(fruitSlices))
	fmt.Println("fruitSlices", fruitSlices[1:3])
	fmt.Println("fruitSlices", fruitSlices[1:1])
	fmt.Println("fruitSlices", fruitSlices[0:1])
	fmt.Println("fruitSlices", fruitSlices[:0])
	fmt.Println("fruitSlices", fruitSlices[:0])
	fmt.Println("fruitSlices", fruitSlices[1:])


}

