package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type RecipePage struct {
	Title, Content string
}

func getFileContent(fileName string, directive string) ([]byte, error) {
	var err error
	var fileContent []byte

	switch directive {
	case "recipes":
		fileContent, err = ioutil.ReadFile("recipes/" + fileName)

	case "root":
		fileContent, err = ioutil.ReadFile(fileName)
	}

	return fileContent, err
}

func parseTemplate(fileName string, fileContent string) []byte {
	var rPage = RecipePage{fileName, fileContent}

	var t = template.Must(template.New("page").ParseFiles("basic_template.txt"))

	var buf bytes.Buffer

	t.ExecuteTemplate(&buf, "basic_template.txt", rPage)

	return buf.Bytes()
}

func prepareResponse(directive string, fileName string) ([]byte, error) {
	var response []byte
	var err error
	var fileContent []byte

	switch directive {

	case "root":
		response, err = getFileContent("index_template.txt", "root")
		
	case "recipes":

		fileContent, err = getFileContent(fileName, "recipes")

        response = []byte(parseTemplate(fileName, string(fileContent)))
	}

	return response, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	var response []byte = []byte("Resource Not Found!")
	
	var err error

	/* Checks if path is at least /recipes, assumes there is more to the URL
	   since it is ensuring that the length is greater than /recipes (8). */
	if ((len(r.URL.Path) > 8) && (r.URL.Path[1:8] == "recipes") ){
		response, err = prepareResponse("recipes", r.URL.Path[8:])

	// Checks if path is root or exactly /recipes, loads homepage.
	} else if (r.URL.Path == "/") || (r.URL.Path[1:8] == "recipes")   {
		response, err = prepareResponse("root", "")
	}

	/* If there is any kind of error preparing the response, 
           handle it by displaying the index page. */
	if (err != nil) {
		response = []byte("ERROR, you've reached an invalid page.")
	}

	//Write response.
	fmt.Fprintf(w, "%s", string(response))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
