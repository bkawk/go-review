package main

import "fmt"

func main() {
	// add up the number in an array
	cups := []float64{1, 2.5, 3, 4.5, 5}
	fmt.Println("Total number of cups needed:", howMany(cups))

	// pass 2 numbers to a single function
	num1, num2 := doubleReturn(5)
	fmt.Println(num1, num2)

	fmt.Println(whatsLeft(1, 2, 3, 4, 5))
}

// accept array of floats as an argument and return float
func howMany(cups []float64) float64 {
	totalCups := 0.0
	for _, amount := range cups {
		totalCups += amount
	}
	return totalCups
}

// return 2 numbers!
func doubleReturn(number int) (int, int) {
	return number * 1, number * 2
}

// we dont know how many arguments are being passed but we know the type
func whatsLeft(cups ...int) int {
	totalCups := 0
	for _, amount := range cups {
		totalCups -= amount
	}
	return totalCups

}

// go run functions.go
