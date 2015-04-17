package main

import(
	"testing"
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"os"
)

/* all functions need to start with the word Test, any other chars can come 
   afterwards provided that the first char is not a lowercase letter. */

// Ping Google to test network connection.
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

/* Make sure server is on. Planning on adding capabilities to start it if not...
   I don't think the http or http/test packages have a way of stopping a go test
   routine if one of the tests fail. If this test fails, there's no need to run 
   the remaining tests. Thus, let tester know that it failed then exit. */
func TestServerIsOnline(t *testing.T) {

	response, err := http.Get("http://golang.recipes")
	if (err != nil) {
		t.Errorf("%v", err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	if (strings.Contains(string(body), "Golang Recipes") != true) {
		t.Errorf("Cannot find page!")
	
		fmt.Println("ERROR: TestServerIsOnline Failed.")	

	os.Exit(1)
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

func TestInvalidPageReachedFomIndex(t *testing.T) {

	var invalidPages = []string{"golang.recipesesofiheofih/",
		"golang.recipes/recipeskjgku/",
		"golang.recipes/recipes/lugluygkuyg",
		"golang.recipes/recipes/sdfihcdso/dfidsfoih"}

	var failedPages = []string{}

	for _, key := range invalidPages {
		response, err := http.Get(key)
		if err != nil {
			t.Errorf("%v", err)
		}
		
		body, _ := ioutil.ReadAll(response.Body)
		
		if (strings.Contains(string(body), "resource not found")) != true {
			failedPages = append(failedPages, key)
		}
	}

	if len(failedPages) != 0 {
		t.Errorf("Failed Pages")
	}

}

