package main

import "fmt"

func main() {
	fmt.Println(sum(42, 13))
	//string-->
	fmt.Println(name("daniel", 14))

}

// the first () its to specify the parameters and its value
// after that you specify the output type and in the {} you specify the function
func sum(x int, y int) int {
	return x + y
}

// if the output has more than 1 type use ()
func name(x string, y int) (string, int) {
	return x, y
}

//other way to make a func is by creating a var ex:

//func sum(x int, y int) int {
//valueSum := x + y
//return valueSum
//}

//and if you create a var with this function its type would be defined by the function type
