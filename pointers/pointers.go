package main

import "fmt"

func main() {
	spoons := 10
	cups := 0
	buckets := new(int)

	changeSpoons(spoons) // the value of sppons gets passed
	changeCups(&cups)    // cups gets passed not the value of cups! note the &
	changeBuckets(buckets)

	fmt.Println("Spoons:", spoons)
	fmt.Println("Cups:", cups)
	fmt.Println("Buckets:", *buckets)
}

func changeCups(cups *int) {
	*cups = 2
}

func changeSpoons(spoons int) {
	spoons = 20
	fmt.Println("Spoons inside changeSpoons", spoons)
}

func changeBuckets(buckets *int) {
	*buckets = 200
}

// go run structs.go
