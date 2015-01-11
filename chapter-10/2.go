/*
http://www.golang-book.com/10/index.htm
Write your own Sleep function using time.After.
*/

package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		return
	}
}

func printNow() {
	fmt.Printf("Current time: %s\n", time.Now())
}

func main() {
	printNow()
	sleep(3)
	printNow()
}
