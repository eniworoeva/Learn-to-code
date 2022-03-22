package main

import "fmt"

//	assign the following values to the three variables
//	a. for x assign 42
//	b. for y assign “James Bond”
//	c. for z assign true

//	use fmt.Sprintf to print all of the VALUES to one single string.
//	ASSIGN the returned value of TYPE string using the short
//	declaration operator to a VARIABLE with the IDENTIFIER “s”
//	print out the value stored by variable “s”

var xen int = 29
var yen string = "Oreva is a Go developer"
var zen bool = true

func main() {
	s := fmt.Sprintf("%v, and he is %v years old. All the above are %v \n", yen, xen, zen)
	fmt.Println(s)

}
