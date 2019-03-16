package main

import "fmt"

func main() {
	fmt.Println(divideThese(3, 0))
	fmt.Println(divideThese(4, 2))
}

func divideThese(num1, num2 int) int {

	defer func() { // defer ensured this runs last
		// fmt.Println(recover()) // recover catches the error and print it
		recover() // doesnt print and doesnt crash
	}()

	answer := num1 / num2
	return answer
}

// go run recover.go
