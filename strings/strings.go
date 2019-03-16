package main

import "fmt"

func main() {

	var hello = "Hello" // string
	var world = `World` // string
	// Concatenation
	fmt.Println(hello, "-", world, "test")      // Joining variables and strings
	fmt.Println(hello, world)                   // Adds a space between two variables
	fmt.Println(hello + " there!")              // Needs a space with +
	fmt.Println("Make a new line \n \n there!") // Makes a new line with \n
	fmt.Println(len(hello))                     // len() gets length

}
