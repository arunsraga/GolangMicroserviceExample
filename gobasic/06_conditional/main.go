package main


import "fmt"

func main(){
	x := 55
	y := 10
	if x<y {
		fmt.Println("%d is smaller than %d", x, y) // %d not replaces
		fmt.Printf("%d is smaller than %d\n", x, y)
	} else {
		fmt.Printf("%d is smaller than %d\n", y, x)
	}

	color := "red1"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "green" {
		fmt.Println("color is green")
	} else {
		fmt.Println("color is "+color)
	}

	
	//Switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "green":
		fmt.Println("color is green")
	default :
		fmt.Println("color is not red or green")



	}


}

