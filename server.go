package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"bufio"
)

type RecipePage struct {
	Title, Content       string
	RecipeName []string
	FileName []string
}

func getSliceOfFileNames(dir string) []string {
	var sliceOfFileInfos []os.FileInfo
	var err error
	var fileNames []string

	sliceOfFileInfos, err = ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, fileInfo := range sliceOfFileInfos {
		fileNames = append(fileNames, fileInfo.Name())
	}

	return fileNames
}

func getSliceOfRecipeNames(dir string) []string {
	fileNames := getSliceOfFileNames(dir)

	var rNames []string

	for _, fname := range fileNames {

		filePointer, err := os.Open(dir + fname)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(filePointer)

		if scanner.Scan() {
			rNames = append(rNames, (scanner.Text())[2:])
		} else {
			panic("EMPTY SCAN")
		}

		filePointer.Close()
	}

	return rNames
}

func (self RecipePage) GenerateLinks() []string {
	var linkSlice []string

	for index, rName := range self.RecipeName {
		link := "<p><a href=\"" + self.FileName[index] + "\">" + rName + "</a></p>"

		linkSlice = append(linkSlice, link)
	}

	return linkSlice
}

func getFileContent(fileName string, directive string) ([]byte, error) {
	var err error
	var fileContent []byte

	switch directive {
	case "recipes":
		fileContent, err = ioutil.ReadFile("recipes/" + fileName)

	case "root": 
	case "template":
		fileContent, err = ioutil.ReadFile(fileName)
	}

	return fileContent, err
}

func parseTemplate(fileName string, fileContent string, directive string) []byte { /*
		var rPage = RecipePage{fileName, fileContent}

		var t = template.Must(template.New("page").ParseFiles("basic_template.txt"))

		var buf bytes.Buffer

		t.ExecuteTemplate(&buf, "basic_template.txt", rPage)
	*/

	var rPage RecipePage
	var t *template.Template
	var buf bytes.Buffer
	var templateFileName []byte

	switch directive {
	case "root":
//		fileNameSlice := getSliceOfFileNames("./recipes/")
//		recipeNameSlice := getSliceOfRecipeNames("./recipes/")

//		rPage = RecipePage{fileName, fileContent, fileNameSlice, recipeNameSlice}
		templateFileName = []byte("index_template.txt")
	case "recipe":
//		rPage = RecipePage{fileName, fileContent, []string{""}, []string{""}}
		templateFileName = []byte("index_template.txt")
	}
		templateFileName = []byte("index_template.txt")

		fileNameSlice := getSliceOfFileNames("./recipes/")
		recipeNameSlice := getSliceOfRecipeNames("./recipes/")

	rPage = RecipePage{fileName, fileContent, fileNameSlice, recipeNameSlice}

	t = template.Must(template.New("page").ParseFiles(string(templateFileName)))

	t.ExecuteTemplate(&buf, string(templateFileName), rPage)

	return buf.Bytes()
}

func prepareResponse(directive string, fileName string) ([]byte, error) {
	var response []byte
	var err error
	var fileContent []byte

	switch directive {

	case "root":
		fileContent, err = getFileContent("index_template.txt", "root")
//		fileContent, err = getFileContent(fileName, "root")

		response = []byte(parseTemplate(fileName, string(fileContent), "root"))
	case "recipes":

//		fileContent, err = getFileContent("basic_template.txt", "template")
		fileContent, err = getFileContent("index_template.txt", "root")

		response = []byte(parseTemplate(fileName, string(fileContent), "recipes"))
	}

	return response, err
}

func handler(w http.ResponseWriter, r *http.Request) {
	var response []byte = []byte("Resource Not Found!")
	var err error

	var pathLength int = len(r.URL.Path)
	var path string = r.URL.Path

	/* Checks if path is at least /recipes, assumes there is more to the URL
	   since it is ensuring that the length is greater than /recipes (8). */
	if (pathLength > 9) && (path[1:8] == "recipes") {
		response, err = prepareResponse("recipes", r.URL.Path[9:])
		// r.URL.Path[8:] = /blah.go
		//response = []byte(r.URL.Path[8:])
		// Checks if path is root or exactly /recipes, loads homepage.
	} else if ((path == "/") || (path == "/recipes") || (path == "/recipes/")){
		response, err = prepareResponse("root", "")
//		response = getFileContent(
	}

	/* If there is any kind of error preparing the response,
	   handle it by displaying the index page. */
	if err != nil {
//		response = []byte("ERROR, you've reached an invalid page." + string(err)) 
//		panic(err)
		response, _ = getFileContent("basic_template.txt", "template")
	}

	//Write response.
	fmt.Fprintf(w, "%s", string(response))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
