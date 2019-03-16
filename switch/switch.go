package main

import "fmt"

func main() {

	height := 120

	switch height {
	case 100:
		fmt.Println("Hey Shorty")
	case 120:
		fmt.Println("Hey Normy")
	case 140:
		fmt.Println("Hey Tallie")
	case 160:
		fmt.Println("Hey Giant")
	}
}

// go run switch.go
