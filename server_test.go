package main

import(
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
//	"fmt"
)

/* all functions need to start with the word Test, any other chars can come 
   afterwards provided that the first char is not a lowercase letter. */
func TestGetRequestSuccess(t *testing.T) {

	response, err := http.Get("http://www.google.com")
	if (err != nil) {
		t.Errorf("%v", err)
	}

	body, _ := ioutil.ReadAll(response.Body)
	
//	fmt.Printf("%s", body)
	
	if (strings.Contains(string(body), "Google Search") != true) {
		t.Errorf("Cannot find page!")
	}
}

func TestServerIsOnline(t *testing.T) {

	response, err := http.Get("http://golang.recipes")
	if (err != nil) {
		t.Errorf("%v", err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	if (strings.Contains(string(body), "Golang Recipes") != true) {
		t.Errorf("Cannot find page!")
	}
}

func TestInvalidPageReachedFromIndex(t *testing.T) {
	
	response, err := http.Get("http://golang.recipes/eofvfduof")
	if (err != nil) {
		t.Errorf("%v", err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	if (strings.Contains(string(body), "resource not found")) != true {
		t.Errorf("Reached valid page with invalid URL.")
	}
}
