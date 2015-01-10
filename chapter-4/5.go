/*
   http://www.golang-book.com/4/index.htm
   Using the example program as a starting point, write a program that converts
   from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
*/

package main

import "fmt"

func main() {
	fmt.Print("Enter a number in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	output := (fahrenheit - 32.0) * (5.0 / 9.0)

	fmt.Println(output)
}

/*
   Enter a number in Fahrenheit: 500
   260
*/
