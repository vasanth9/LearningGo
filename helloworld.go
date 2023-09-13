package main

import "fmt"

func main() {
	fmt.Print("Hello World")
	fmt.Println("Hello World")
	var a int
	var b string
	fmt.Println("Before assigning values AB = ", a, b)
	a = 10
	b = "Whatsapp"
	fmt.Println("After assigning values AB = ", a, b)

	var (
		x int
		y bool
	)
	fmt.Println("Before assigning values XY = ", x, y)
	i, j := 100, true
	fmt.Println("After assigning values XY = ", i, j)
}
