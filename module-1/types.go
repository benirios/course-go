package main

import (
	"fmt"
)

func main() {

	//bool (t/f)
	fmt.Printf("Type: %T - Value:%v", true, true)
	//string (bytes sequence)
	fmt.Printf("Type: %T - Value:%v", "string", "string")
	//int
	fmt.Printf("Type: %T - Value:%v", 1, 1)
	//float (64/32)
	fmt.Printf("Type: %T - Value:%v", 1.1, 1.1)

}
