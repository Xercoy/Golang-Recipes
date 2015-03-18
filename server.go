package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getFileContent(fileName string) string {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fileContent = []byte("Error: resource not found!")
	}

	return string(fileContent)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var fileName string = string(r.URL.Path[1:])
	fileContent := getFileContent(fileName)

	//Write response
	fmt.Fprintf(w, string(fileContent))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
