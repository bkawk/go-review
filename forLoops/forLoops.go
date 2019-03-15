package main

import "fmt"

func main() {

	// Increment
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Decrement
	x := 10
	for x >= 1 {
		fmt.Println(x)
		x--
	}

	// same as above but formatted diferently
	for y := 0; y < 5; y++ {
		fmt.Println(y)
	}

	// for range
	var fruits = [4]string{"Apple", "Orange", "Pear", "Watermelon"}
	for j, value := range fruits {
		fmt.Println(value, j)
	}

}

// go run forLoops.go
