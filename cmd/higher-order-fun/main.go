package main

import "fmt"

func double(a int, b int, action func(int, int) int) int {
	return 2 * action(a, b)
}

func main() {
	sum := double(
		10,
		20,
		func(a int, b int) int { return a * b },
	)
	fmt.Println(sum)
}
