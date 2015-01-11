/*
http://www.golang-book.com/8/index.htm
Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should
give you x=2 and y=1).
*/

package main

import "fmt"

func swap(x int, y int) (int, int) {
	x ^= y
	y ^= x
	x ^= y
	return x, y
}

func main() {
	fmt.Println("Enter two numbers to swap: ")
	var x, y int
	fmt.Scanf("%d,%d", &x, &y)
	fmt.Println(swap(x, y))
}
