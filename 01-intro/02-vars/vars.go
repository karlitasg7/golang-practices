package main

import "fmt"

func main() {
	var name string
	name = "Karlitas"

	age := 20

	fmt.Println(name, age)

	// default values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// constants
	const pi = 3.141592
	fmt.Println(pi)
}
