package main

import "fmt"

func main() {
	defer atEnd() // will always run last
	stepOne()
}

func stepOne() { fmt.Println(1) }
func atEnd()   { fmt.Println(2) }

// go run defer.go
