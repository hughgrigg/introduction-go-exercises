/*
Write another program that converts from feet into meters. (1 ft = 0.3048 m)
*/

package main

import "fmt"

func main() {
	fmt.Println("Enter a distance in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048

	fmt.Printf("%.2f feet is %.2f meters\n", feet, meters)
}
