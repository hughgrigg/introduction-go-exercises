/*
http://www.golang-book.com/7/index.htm
Write a function which takes an integer and halves it and returns true if it was
even or false if it was odd. For example half(1) should return (0, false) and
half(2) should return (1, true).
*/

package main

import "fmt"

func halfAndEven(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	fmt.Println("Enter an integer: ")
	var x int
	fmt.Scanf("%d", &x)
	half, isEven := halfAndEven(x)
	fmt.Printf("Half of %d is %.1f\n", x, half)
	fmt.Printf("%d is even? %t\n", x, isEven)
}
