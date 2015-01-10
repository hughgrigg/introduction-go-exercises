/*
http://www.golang-book.com/7/index.htm
The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1)
+ fib(n-2). Write a recursive function which can find fib(n).
*/

package main

import "fmt"

func getFibonacciCalculator() func(n int) uint64 {
	fibonaccis := map[int]uint64{0: 0, 1: 1}
	var fib func(n int) uint64
	fib = func(n int) uint64 {
		if n < len(fibonaccis) {
			return fibonaccis[n]
		}
		fibonaccis[n] = fib(n-1) + fib(n-2)
		return fibonaccis[n]
	}
	return fib
}

func main() {
	fib := getFibonacciCalculator()
	fmt.Println("Enter fibonacci index: ")
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(fib(n))
}

/*
This was my first solution, using a recursive closure. The solution at this link
is more idiomatic Go, however:
https://codereview.stackexchange.com/questions/28386/fibonacci-generator-with-golang/28445#28445

My solution here also fails where n > 93, as the 94th fibonacci number is
19740274219868223167, higher than uint64's max value of
18446744073709551615
*/
