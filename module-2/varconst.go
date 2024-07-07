package main

import (
	"fmt"
)

func main() {
	//variables
	var name string
	name = "breno"
	fmt.Println(name)
	//It prints both if the value changes
	name = "beni"
	fmt.Println(name)
	//It dont acceps any other type of value if its not the specified
	// var yo int
	// yo = "bento"
	// fmt.Println(yo)
	var age int
	age = 1
	fmt.Println(age)
	// other way using type inference
	// if its not a string use the type after the var name
	var a = "beni"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b)
	fmt.Println(c)
	//bool
	var be = true
	fmt.Println(be)
	// it uses the type inference
	// to create the var without the need of specify the var type
	f := "breno"
	fmt.Println(f)

	//Consts
	// Its a value that dont change
	const idadeBeni = 15
	fmt.Println(idadeBeni)
}
