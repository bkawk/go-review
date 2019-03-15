package main

import "fmt"

func main() {
	var age int = 40 // int
	// float
	var floatNum float64 = 1.6180339 // float
	// uint
	var max255 uint8 = 255                                    // 0 = 255
	var max65535 uint16 = 65535                               // 0 - 65535
	var max4294967295 uint32 = 4294967295                     // 0 - 4294967295
	var max18446744073709551615 uint64 = 18446744073709551615 // 0 - 18446744073709551615
	// int
	var min128 int8 = -128                                  // -128 to 127
	var min32768 int16 = -32768                             // -32768 to 32767
	var min2147483648 int32 = -2147483648                   // -2147483648 to 2147483647
	var min9223372036854775808 int64 = -9223372036854775808 // -9223372036854775808 to 9223372036854775807

	fmt.Println(age)
	fmt.Println(floatNum)
	fmt.Println(max255)
	fmt.Println(max65535)
	fmt.Println(max4294967295)
	fmt.Println(max18446744073709551615)
	fmt.Println(min128)
	fmt.Println(min32768)
	fmt.Println(min2147483648)
	fmt.Println(min9223372036854775808)
	fmt.Printf("%.2f \n", floatNum) // float formating to 2 decimal places

}
