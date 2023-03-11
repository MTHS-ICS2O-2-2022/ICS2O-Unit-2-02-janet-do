// Created by: Janet Do
// Created on: Sep 2020
//
// This program does basic math
package main

import (
	"fmt"
	"math"
)

func main() {
	// This function does basic math
	fmt.Println("9 + 2 = ", (9 + 2))
	fmt.Println("7 - 3 = ", (7 - 3))
	fmt.Println("4 * 2 = ", (4 * 2))
	fmt.Println("4 + 4 / 2 = ", (4 + (4 / 2)))
	fmt.Println("5 + 2³ = ", (5 + math.Pow(2, 3)))

	fmt.Println("\nDone.")
}
