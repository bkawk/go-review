package main

import "fmt"

func main() {

	// Formating
	fmt.Printf("%.2f \n", 1.23456789) // float formating to 2 decimal places
	fmt.Printf("%T \n", 1.2345)       // get data type
	fmt.Printf("%t \n", true)         // get bool type
	fmt.Printf("%d \n", 100)          // get int type
	fmt.Printf("%b \n", 90)           // get binary
	fmt.Printf("%c \n", 80)           // get data type
	fmt.Printf("%x \n", 'A')          // get hex type
	fmt.Printf("%e \n", 100.123)      // get scientific type
}

// go run formatting.go
