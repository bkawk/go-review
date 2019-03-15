package main

import "fmt"

func main() {

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Pear"
	fruits[3] = "Watermelon"

	fmt.Println(fruits)
	fmt.Println(fruits[0])

	// prints 1 to 3 indexes
	fmt.Println(fruits[1:3])
	// prints 0 to 3 indexes
	fmt.Println(fruits[:3])
	// prints 0 to 3 indexes
	fmt.Println(fruits[0:3])
	// prints index 3 to the end
	fmt.Println(fruits[3:])

	// Same as above just formatted diferently
	var fruitsAgain = [4]string{"Apple", "Orange", "Pear", "Watermelon"}
	fmt.Println(fruitsAgain)
	fmt.Println(fruitsAgain[0])

	// length 4 and fills 4
	var lengths = [4]float64{1, 2, 3, 4}
	fmt.Println(lengths)
	fmt.Println(lengths[0])

	// length 6 but fills 4, gives two 0's
	var widths = [6]float64{1, 2, 3, 4}
	fmt.Println(widths)
	fmt.Println(widths[0])

	// length unspecified but fills 4
	var heights = []float64{1, 2, 3, 4}
	fmt.Println(heights)
	fmt.Println(heights[0])
}

// go run arrays.go
