package main

import "fmt"

func area(radius float64) float64 {
	return 22 / 7 * radius * radius
}

func circumference(radius float64) float64 {
	return 2 * 22 / 7 * radius
}

func main() {
	fmt.Println(area(7))
	fmt.Println(circumference(7))
}
