package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("some-text.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("This is some sample text")
	file.Close()

	stream, err := ioutil.ReadFile("some-text.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}

// go run fileio.go
