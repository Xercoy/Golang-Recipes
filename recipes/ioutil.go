//ioutil
package main

import (
	"fmt"
	"io/ioutil" //Info: http://golang.org/pkg/io/ioutil/
	"os"
)

func main() {

	/* Readfile Example Begin */
	var sliceOfFileInfos []os.FileInfo
	var err error

	// CHANGE AS NEEDED
	var dirName string = "./"

	/* ioutil.ReadDIr returns slice of os.FileInfos, a data
	   structure from the os package from a specified directory.
	   which contains useful information about a file, such as
	   name, size, time modified, and a few more details.
	   For more info: http://golang.org/pkg/os/#FileInfo */
	sliceOfFileInfos, err = ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	/* Iterate through slice and print the Name field,
	   essentially the name of the file. */
	fmt.Printf("\nList of files in directory:\n")
	for _, fileInfo := range sliceOfFileInfos {
		fmt.Printf("%s\n", fileInfo.Name())
	}
	/* Readfile Example End */
}
