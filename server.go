package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
        var fileName string = string(r.URL.Path[1:])

        file, err := ioutil.ReadFile(fileName)		
	if (err != nil) {
		fmt.Fprintf(w, "Error 404: resource not found!")
	}

	fmt.Fprintf(w, string(file))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
