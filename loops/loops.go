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

	// for range with unused iterator
	var cups = []float64{1, 2.5, 3, 4.5, 5}
	totalCups := 0.0
	for _, amount := range cups {
		totalCups += amount
	}
	fmt.Println(totalCups)

}

// go run forLoops.go
