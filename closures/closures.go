package main

import "fmt"

// Outer function
func main() {
	cups := 3

	// Innrer function
	replicator := func() int {
		cups *= 2
		return cups
	}

	fmt.Println(replicator()) // 6
	fmt.Println(replicator()) // 12
	fmt.Println(replicator()) // 24
}

// go run closures.go
