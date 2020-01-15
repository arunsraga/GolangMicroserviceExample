package main


import "fmt"

func main(){
	// Pointers : poiint to the memory address/location of the value
	a := 5 
	fmt.Println("a mem address is is ", &a)
	b := &a

	fmt.Println("b mem address is ", b)

	*b = 20

	fmt.Println("a is ", a)
	fmt.Println("b is ", *b)
}

