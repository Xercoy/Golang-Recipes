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

// Helper function, replaces 3 lines of error checking with 1.
func panicIfNotNil(e error) {
	if e != nil {
		panic(e)
	}
}

// Ping Google to test network connection.
func TestGetRequestSuccess(t *testing.T) {

	response, err := http.Get("http://www.google.com")
	if (err != nil) {
		t.Errorf("%v", err)
	}

	body, _ := ioutil.ReadAll(response.Body)

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

func TestInvalidIndexPages(t *testing.T) {
	/* Two responses, one that is definitely bad, and another that might be 
           bad. Since an error page may change over time, compare the two bodies
           instead of hardcoding anything. */
	var testUrlResp* http.Response 
	var posBadUrlResp* http.Response
	var err error
	var pagesReached []string // If we get through to any, keep track of it.

	// Urls being tested
	var testUrls = []string{"http://golang.recipes/recipesxxx/", // with '/'
		"http://golang.recipes/recipes/xxxxxx/xxxx"} // after two levels

	/* Get resp + body for a page that is definitely invalid to compare. */
	posBadUrlResp, err = http.Get("http://golang.recipes/xxxxxx")
	panicIfNotNil(err)
	posBadUrlBody, err := ioutil.ReadAll(posBadUrlResp.Body)
	panicIfNotNil(err)

	/* Iterate through test urls to determine whether they can be reached. 
           If so, append to the slice pagesReached so that tester can be 
           notified. */
	for _, url := range testUrls {
		testUrlResp, err = http.Get(url)
		panicIfNotNil(err)

		testUrlBody, err := ioutil.ReadAll(testUrlResp.Body)
		// A response body will never be nil, even on an empty request.
		if len(testUrlBody) == 0 {
			panic(err)
		}

		// Compare body of an invalid page to the url being tested.
		if string(testUrlBody) == string(posBadUrlBody) {
			pagesReached = append(pagesReached, url)
		}

		testUrlResp.Body.Close()
	}

	if pagesReached != nil {
		t.Errorf("\nPages reached: %v\n", pagesReached)
	}
}
