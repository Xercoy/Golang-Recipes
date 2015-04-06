// Command Line Arguments

package main

import (
	"fmt"
	"os"      // http://golang.org/pkg/os/#pkg-variables
	"strconv" // http://golang.org/pkg/strconv/#ParseInt
)

/* The Args variable in the os package contains a string slice of the command
   line arguments. Since the arguments are all strings, you would need to
   convert them to the desired data type if necessary. The strconv package takes
   care of this. */
func main() {

	var sum int

	// We start at index 1 because 0 contains the program name.
	for i := 1; i < len(os.Args); i++ {

		num, err := strconv.Atoi(os.Args[i])
		if err != nil {
			panic(err)
		}

		sum += num
	}

	fmt.Printf("\nThe sum of command line arguments is %d.\n\n", sum)
}

/* To convert to different integer types use types use:
                   strconv.ParseInt(os.Args[i], 10, 0)
                   This is equivalent to strconv.atoi, except it returns int64.
                   Convert the string to an integer, strconv.ParseInt receives
		   the string to be converted, the base it should be converted
		   to, and the integer size that it shoudl fit into (0, 8, 16,
		   32, 64).*/
