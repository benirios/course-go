//LISTS
//1-arrays and slices will always be the same type
//[1,2,3,4,5,6,7,8,9] - []int
//[breno,beni,celaine,sarah] - []string
//[1.1,2.2,3.3,4.4,5.5,6.6,7.7,8.8,9.9] - []float64

// maps - can mix different value types
// we have a structure "key - value"
// the key have a type and the value can have a different type
// map[string]int
// {"beni": 15, "breno": 29}

//Array
//-fixed length
//-access the values through indexes: a[0], a[1]...
//-the function len() returns the array length
//-func append() add values to the array

package main

import (
	"fmt"
)

func main() {
	//array
	var array [2]string
	array[0] = "hello"
	array[1] = "world"
	//print the specified index
	fmt.Println(array[0], array[1])
	//prints the list
	fmt.Println(array)

	//Remember that the array can be declared like a var using the type inference ex:
	numPrims := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrims)
	fmt.Println(numPrims[0:3])
	fmt.Println(numPrims[:3])

	//Slices
	//var slice []string
	slice := make([]string, 2)
	slice[0] = "hi"
	slice[1] = "everyone"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	// it dont have a defined length to base on it
	//to solve that we use the 'make' to "define" a length to it like in the line 43

	//len specific function
	fmt.Println(len(slice))
	//append to add values
	slice = append(slice, "beni")
	fmt.Println(slice)
}
