package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {
	A := person{name: "Harikrishnan", age: 27, city: "Bangalore"}
	fmt.Println(A)
}
