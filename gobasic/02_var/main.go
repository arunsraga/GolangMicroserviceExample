package main


import "fmt"

// test := "abc"
// for above code initialization outside not alowed
// non-declaration statement outside function body

func main(){
	// fmt.Println("Hello World")
	// Data Types 
	// string
	// bool
	// int 
	// int int8 int16 int32 int64
	// uint uint8 unit16 unit32 unit64 unitptr
	// byte - alias for unit8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	
	// Usin var 
	var name string = "abcd"
	
	// declaring and not using varaibale result into error 
	var i int = 10 
	// 02_main/main.go:23:6: i declared and not used

	
	test := "abc"

	fmt.Println("test", test)// no semicolon

	fmt.Println("name", name)// no semicolon
	fmt.Println("int i", i)// no semicolon


	name, email := "abcd", "abcd@mail.com"

	fmt.Println("name and email", name, email)

}

