package main

import "fmt"

func main() {

	height := 120

	if height == 120 {
		fmt.Println("You are 120cm high")
	}

	if height <= 200 {
		fmt.Println("You are under 200cm high")
	}

	if height >= 100 {
		fmt.Println("You are taller than 100cm")
	}

	if height >= 300 {
		fmt.Println("You are taller than 100cm")
	} else {
		fmt.Println("You are not taller than 300cm")
	}

	if height >= 300 {
		fmt.Println("You are taller than 100cm")
	} else if height >= 200 {
		fmt.Println("You are taller than 200cm")
	} else {
		fmt.Println("You are not taller than 200cm")
	}
}

// go run if.go
