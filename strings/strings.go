package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var hello = "Hello" // string
	var world = "World" // string
	var csv = "1,2,3,4,5,6,7"
	var fruits = []string{"Apple", "Orange", "Pear", "Watermelon"}
	var listOfFruits = strings.Join(fruits, ",")     // join on ","
	fmt.Println(hello, "-", world, "test")           // Joining variables and strings
	fmt.Println(hello, world)                        // Adds a space between two variables
	fmt.Println(hello + " there!")                   // Needs a space with +
	fmt.Println("Make a new line \n \n there!")      // Makes a new line with \n
	fmt.Println(len(hello))                          // len() gets length
	fmt.Println(strings.Contains(hello, "lo"))       // hello contains lo true
	fmt.Println(strings.Index(hello, "lo"))          //index of lo is 3
	fmt.Println(strings.Count(hello, "l"))           // there are 2 l's in hello
	fmt.Println(strings.Replace(hello, "l", "x", 1)) // replace te first l with an x
	fmt.Println(strings.Split(csv, ",")[3])          // split onthe comma into a list and then show index 3
	sort.Strings(fruits)                             // sort the fruits
	fmt.Println("Fruits in order:", fruits)          // print out the ordered fruits
	fmt.Println("List of fruits", listOfFruits)      // print out the list of fruits

}

// go run strings.go
