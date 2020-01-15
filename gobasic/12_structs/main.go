package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	age int
	gender string
}

//greeting method e.g of value rxr
func(p Person) greet() string {
	return "Hello, My name is "+p.firstName 
} 

//update age e.g of pointer rxr

func(p *Person) updateAge()  {
	p.age = 3000
} 

func main(){
	// init  person using struct
	person1 := Person{firstName: "abcd", lastName:"xyz", city: "test1", age: 32, gender:"ffff"  }

	// or
	person2 := Person{ "abcd1", "xyz1", "test12", 132, "df"  }
	


	fmt.Println("person1 struct is ", person1)
	fmt.Println("person2 struct is ", person2)

	// Value Reciever 
		
	fmt.Println(" Greeting :",person1.greet())
	fmt.Println(" person1 Before :",person1)

	person1.updateAge()

	fmt.Println(" person1 after ageUpdate :", person1)

}

