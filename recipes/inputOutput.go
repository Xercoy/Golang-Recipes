//User Input and Output

package main

import (
	"fmt" // https://golang.org/pkg/fmt/
)

func main() {
	/* fmt.Scanf example begin */
	var name string
	var age int

	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)

	fmt.Println("How old are you?")
	fmt.Scanf("%d", &age)

	if age%2 == 0 {
		fmt.Printf("Hello, %s! Your age, %d, is an even number!\n\n", name, age)
	} else {
		fmt.Printf("Hello, %s! Your age, %d, is not an even number!\n\n", name, age)
	}
	/* fmt.Scanf example end */
}
