/*
http://www.golang-book.com/7/index.htm
Using makeEvenGenerator as an example, write a makeOddGenerator function that
generates odd numbers.
*/

package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	for i := 0; i < 20; i++ {
		fmt.Println(nextOdd())
	}
}
