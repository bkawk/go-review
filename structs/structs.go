package main

import "fmt"

func main() {

	myCar := vehicle{4, 8, 2}
	myBike := vehicle{2, 1, 2}
	fmt.Println("My cars has", myCar.wheels, "wheels")
	fmt.Println("My bike has", myBike.wheels, "wheels")

}

type vehicle struct {
	wheels  int
	lights  int
	mirrors int
}

// go run structs.go
