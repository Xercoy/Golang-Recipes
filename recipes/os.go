//os
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* Open Example Begin */
	/* This example opens a file, reads it one line at a time,
	   and prints the line to stdout directly after it is read. */

	var filePointer *os.File
	var err error
	var scanner *bufio.Scanner

	// CHANGE AS NEEDED
	var fileName string = "ioutil.go"

	filePointer, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner = bufio.NewScanner(filePointer)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	filePointer.Close()
	/* Open Example End */
}
