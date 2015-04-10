package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
//	"strings"
)

//const recipesPath string = "./recipes/" 
const recipesPath string = "recipes"

/* Considered creating diff structs for each directive but that means more
   coding and more things to keep track of... Who knows, maybe there will be a
   need for a page to keep track of all of this info. Maybe the master struct
   should be an aggregation of all of the different structs? */
type RecipePage struct {
	Title, Content string
	RecipeName     []string
	FileName       []string
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

// Appending a string: http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
func (self RecipePage) GenerateLinks() template.HTML {
	var buffer bytes.Buffer

	for index, rName := range self.RecipeName {
		link := "<p><a href=\"/recipes/" + self.FileName[index] + "\">" + rName + "</a></p>"

		buffer.WriteString(link)
	}

	return template.HTML(buffer.String())
}

/* Return the contents of a file as a slice of bytes. The directive tells the
   function where to look for the files. For instance, all recipes reside in the
   recipes folder. Need to place templates in their own folder... */
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

/* parseTemplate is called by responseHandler. This function receives the
   information needed such as file names so that it could run said data
   against the template and return the result. */
func parseTemplate(fileName string, fileContent string, directive string) []byte {
	var rPage RecipePage
	var t *template.Template
	var buf bytes.Buffer
	var templateFileName []byte

	switch directive {
	case "root":
		// Get a slice of the recipes' title and file names.
		fileNameSlice := getSliceOfFileNames("./recipes/")
		recipeNameSlice := getSliceOfRecipeNames("./recipes/")

		/* Fill out the data structure accordingly. This template
		   doesn't use the first two fields bc they're for recipes */
		rPage = RecipePage{"", "", recipeNameSlice, fileNameSlice}

	case "recipes":
		/* Recipe template doesn't need every recipes' title/file name
		   so leave both blank */

		// Template for the recipes page.
		rPage = RecipePage{fileName, fileContent, []string{""}, []string{""}}
	}

	templateFileName = []byte(directive + ".txt")

	t = template.Must(template.New("page").ParseFiles(string(templateFileName)))

	t.ExecuteTemplate(&buf, string(templateFileName), rPage)

	return buf.Bytes()
}

/* Based on the directive given, responseHandler calls parseTemplate so that
   the final, complete response */
func responseHandler(directive string, fileName string) ([]byte, error) {
	var response []byte
	var err error
	var fileContent []byte

	switch directive {
	case "root":
		// Generate the template, no content to be filled out here.
		response = []byte(parseTemplate(fileName, string(""), "root"))

	case "recipes":
		// Get content of the recipe.
		fileContent, err = getFileContent(fileName, "recipes")

		// Fill the recipes template with code.
		response = []byte(parseTemplate(fileName, string(fileContent), "recipes"))
	}

	return response, err
}

/* The most important thing pathHandler does is determine whether the resource
   is valid. Depending on whether the desired resource is the root of the page
   or a recipe, pathHandler will call responseHandler to prepare a response
   to the request received. A directive is included to responseHandler as the
   first argument. Since the responses are prepared differently, this tells the
   function whether it is receiving a path for a recipe or for the root page.
   There is no functionality to display an invalid resource at the moment. */
func pathHandler(w http.ResponseWriter, r *http.Request) {
	var response []byte = []byte("Resource Not Found!")
	var err error

	var pathLength int = len(r.URL.Path)
	var path string = r.URL.Path

	/* Checks if path is at least /recipes/, assumes there is more to the URL
	   since it is ensuring that the length is greater than /recipes/ (9). */
	if (pathLength > 9) && (path[:9] == "/recipes/") {

		response, err = responseHandler("recipes", r.URL.Path[9:])

	} else if (path == "/") || 
		path == ("/" + recipesPath) || // /recipes 
		path == ("/" + recipesPath + "/") { // /recipes/

		response, err = responseHandler("root", "")
	}

	/* If there is any kind of error preparing the response,
	   handle it by displaying an error message. */
	if err != nil {
		//panic(err)
	}

	//Write response.
	fmt.Fprintf(w, "%s", string(response))
}

// The pathHandler function is responsible to handling all requests.
func main() {
	http.HandleFunc("/", pathHandler)
	http.ListenAndServe(":8080", nil)
}
