/*
http://www.golang-book.com/7/index.htm
Write a function with one variadic parameter that finds the greatest number in a
list of numbers.
*/

package main

import "fmt"

func largestOf(args ...int) int {
	largest := args[0]
	for _, v := range args {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(largestOf(12, 354, 78, 17, 8, 45, 93))
}
