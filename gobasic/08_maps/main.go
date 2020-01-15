package main

import "fmt"

func main(){
	//maps : is a key value pair
	
	emails := make(map[string]string) // map[key type]value type

	// Assign Key value

	emails["abcd"] = "abcd@mail.com"
	emails["bcd"] = "bcd@mail.com"
	emails["xbcd"] = "xbcd@mail.com"

	fmt.Println(emails)
	fmt.Println(emails["abcd"])

	// delete from map
	delete(emails, "abcd")
	fmt.Println(emails)

	// Declare and assign Key value 
	numbers := map[string]int{"a":12346, "b":8939284293, "c":1312323}
	fmt.Println(numbers)


}

