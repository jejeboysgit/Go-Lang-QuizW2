package main

import (
	"fmt"
	"math"
)

func main() {
	var t, result float64
	fmt.Scanln(&t)
	//result = s(t)
	result = math.Pow(t, 3) - (12 * math.Pow(t, 2)) + (36 * t) - 30
	//output
	fmt.Println(result)
}
