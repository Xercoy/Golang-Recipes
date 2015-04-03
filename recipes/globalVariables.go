//Global Variables
package main

import "fmt"

/* Global variable myNumber. This example creates three myNumber variables
   with different values at different scopes. One is global, the other is local
   to main, and the last is local to a for loop. Observe how they exist in their
   own scope given the result of printing the value in myNumber. */
var myNumber int = 123

func globalMyNumberPrint() {
	fmt.Printf("\nmyNumber from a function outside of main: %d.\n", myNumber)
}

func main() {
	/* myNumber local to main */
	var myNumber int = 456

	/* will output 456 */
	fmt.Printf("\nmyNumber before the for loop: %d.\n", myNumber)

	for true {
		/* myNumber local to the for loop */
		var myNumber int = 789

		/* will output 789 */
		fmt.Printf("\nmyNumber from inside the for loop: %d\n", myNumber)
		break
	}

	/* will output 456 */
	fmt.Printf("\nmyNumber outside of the for loop: %d\n", myNumber)

	/* will output 123 */
	globalMyNumberPrint()
}
