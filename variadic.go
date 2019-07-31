package main

import "fmt"

func sum(nums ...float64) {
	fmt.Print(nums, " ")
	total := 0.0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	const pi = 3.1459
	const two = 2
	a := 12.29
	b := 38.56
	c := 5.23

	sum(pi, two)
	sum(a, b, c)
	sum(pi, a, b)
}
