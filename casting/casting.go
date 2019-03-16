package main

import (
	"fmt"
	"strconv"
)

func main() {
	anInt := 23
	aFloat := 40.23
	aString := "102"
	anotherString := "102.456"

	fmt.Println(float64(anInt))
	fmt.Println(int(aFloat))

	newInt, _ := strconv.ParseInt(aString, 0, 64)
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(anotherString, 64)
	fmt.Println(newFloat)
}

// go run casting.go
