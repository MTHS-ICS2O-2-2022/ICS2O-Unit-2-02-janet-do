// Created by: Janet Do
// Created on: Sep 2020
//
// This program does Area and Perimeter
package main

import "fmt"

func main() {
	length := 5.0
	width := 3.0

	area := length * width
	perimeter := 2 * (length + width)

	fmt.Printf("Area: %.2f cm²\n", area)
	fmt.Printf("Perimeter: %.2f cm\n", perimeter)
}
