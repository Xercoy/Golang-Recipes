package main

import (
//	"bytes"
	"fmt"
//	"html/template"
	"io/ioutil"
	"net/http"
)

type Recipe struct {
	Title, Content string
}
/*
func prepareResponse(title string, content string) string {
	var newRecipe = Recipe{title, content}

	var t = template.Must(template.New("recipe").ParseFiles("basic_template.txt"))

	var buf bytes.Buffer

	t.ExecuteTemplate(&buf, "basic_template.txt", newRecipe)

	return buf.String()
}*/

func getFileContent(fileName string)  string {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContent)
}
/*
func handler(w http.ResponseWriter, r *http.Request) {
	//Retrieve file content.
	var fileName string = string(r.URL.Path[1:])
	fileContent := getFileContent(fileName)

	//Get response ready
	response := prepareResponse(fileName, fileContent)

	//Write response.
	fmt.Fprintf(w, response)
}*/
func prepareResponse(directive string, path string) []byte {
	var response []byte
	
	switch directive {

	case "root":
		response = []byte(getFileContent("index_template.txt"))
	}
	
	return response
}

func handler (w http.ResponseWriter, r *http.Request) {

	var response []byte

	if (r.URL.Path[1:] == "") {
		response =  prepareResponse("root", "")
	} else if ( r.URL.Path[2:8] == "recipes" ) {
		response = prepareResponse("recipes", r.URL.Path[9:])
	}

	//Write response.
	fmt.Fprintf(w, string(response))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


