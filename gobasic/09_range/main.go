package main


import "fmt"

func main(){
	ids := []int{33,43,565,6,7,8}

	// Loop throough ids
	for i, id := range ids {
		fmt.Printf("%d - ID: value is %d\n", i, id)
	}

	// i is not using case
	for _, id := range ids {
		fmt.Printf("ID: value is %d\n", id)
	}

	// Range with map
	numbers := map[string]int{"a":12346, "b":8939284293, "c":1312323}

	for k,v := range numbers {
		fmt.Printf("\n%s - Number: value is %d\n", k, v)
	}


	for _,v := range numbers {
		fmt.Printf(" Number: value is %d\n", v)
	}

}

