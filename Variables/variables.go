package main

import "fmt"

func main() {

	var hello = "Hello"            // string
	var world = `Hello World`      // string
	var age int = 40               // int
	var favNum float64 = 1.6180339 // float
	// uint
	var max255 uint8 = 255                                    // 0 = 255
	var max65535 uint16 = 65535                               // 0 - 65535
	var max4294967295 uint32 = 4294967295                     // 0 - 4294967295
	var max18446744073709551615 uint64 = 18446744073709551615 // 0 - 18446744073709551615
	// int
	var min128 int8 = -128                                  // -128 to 127
	var min32768 int16 = -32768                             // -32768 to 32767
	var min2147483648 int32 = -2147483648                   // -2147483648 to 2147483647
	var min9223372036854775808 int64 = -9223372036854775808 // -9223372036854775808 to 9223372036854775807
	// Concatenation
	fmt.Println(hello, "-", age, "test")        // Joining variables and strings
	fmt.Println(hello, age)                     // Adds a space between two variables
	fmt.Println(hello + " there!")              // Needs a space with +
	fmt.Println("Make a new line \n \n there!") // Makes a new line with \n
	// Arithmatic
	fmt.Println("plus", 1+2)
	fmt.Println("minus", 2-1)
	fmt.Println("times", 2*2)
	fmt.Println("division", 6/2)
	fmt.Println("modulas", 6%4)
	// Booleans
	var isBoring bool = false
	var imHappy bool = true
	// Formating
	fmt.Printf("%.2f \n", favNum) // float formating to 2 decimal places
	fmt.Printf("%T \n", favNum)   // get data type
	fmt.Printf("%t \n", imHappy)  // get data type

	fmt.Println(hello)
	fmt.Println(world)
	fmt.Println(age)
	fmt.Println(favNum)
	fmt.Println(max255)
	fmt.Println(max65535)
	fmt.Println(max4294967295)
	fmt.Println(max18446744073709551615)
	fmt.Println(min128)
	fmt.Println(min32768)
	fmt.Println(min2147483648)
	fmt.Println(min9223372036854775808)
	// Length
	fmt.Println(len(hello))
	fmt.Println(isBoring)
	fmt.Println(imHappy)

}
