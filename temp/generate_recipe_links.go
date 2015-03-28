package main

import (
	"fmt"
//	"text/template"
	"io/ioutil" //http://golang.org/pkg/io/ioutil/
	"os" //https://golang.org/pkg/os/
)

func main() {
	var sliceOfFileInfos []os.FileInfo
	var err error

	sliceOfFileInfos, err = ioutil.ReadDir("./")
	if (err != nil) {
		panic(err)
	}

	for _, fileInfo := range sliceOfFileInfos {
		fmt.Println(fileInfo.Name())
	}
}
