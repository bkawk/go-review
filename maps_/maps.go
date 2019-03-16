package main

import "fmt"

func main() {

	carsAge := make(map[string]int)
	carsAge["Volvo"] = 12
	carsAge["Porsche"] = 4
	carsAge["Lambo"] = 2

	fmt.Println(carsAge["Volvo"])
	fmt.Println(len(carsAge))
	delete(carsAge, "Lambo")
	fmt.Println(len(carsAge))

}

// go run maps.go
