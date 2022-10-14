/**
 * @author Jejeboys
 */
package main

import (
	"fmt"
	"math"
)

func pow(a float64, b float64) float64 {
	return math.Pow(a, b)
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func s(t float64) float64 {
	var a, b, c float64

	a = pow(t, 3)
	b = pow(t, 2)
	b = multiply(b, 12)
	c = multiply(36, t)

	return a - b + c - 30
}

func main() {
	// declaration
	var t, result float64

	// input
	fmt.Print("Enter text: \n")
	fmt.Scanln(&t)

	// process
	result = s(t)

	// output
	fmt.Println(result)
}
