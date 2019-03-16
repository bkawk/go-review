package main

import "fmt"

func main() {
	panicThrow()
}

func panicThrow() {

	defer func() {
		fmt.Println(recover())
	}()
	panic("THROW / PANIC !!") // Name the error to be printed
}

// go run panic.go
