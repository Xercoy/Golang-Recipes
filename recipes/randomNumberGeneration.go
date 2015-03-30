// Generating Random Numbers

package main

import (
	"fmt"
	"math/rand" // http://golang.org/pkg/math/rand/
	"time"
)

func main() {

	/* Specify the seed to the default Source, if it isn't called then he Source will seed with 1. Here we used time.Now so that the seed is different every time the program is run. t.Unix returns number of milaseconds since 1/1/1970*/
	rand.Seed(time.Now().Unix())

	/* The argument to Intn specifies the upper limit of numbers that can be generated, in this example it's 0 - 100 */
	for i := 0; i < 9; i++ {
		fmt.Println(rand.Intn(100))
	}
}
