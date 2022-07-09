package main

import "fmt"

/*Use var to DECLARE three VARIABLES. The variables
should have package level scope. Do not assign
VALUES to the variables. Use the following IDENTIFIERS
for the variables and make sure the variables are of
the following TYPE (meaning they can store VALUES of that TYPE)

print out the values for each identifier
*/

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v,\n", x)
	fmt.Printf("%v, \n", y)
	fmt.Printf("%v, \n", z)
}
