package main

import (
//	"fmt"
	"text/template"
	"io/ioutil" //http://golang.org/pkg/io/ioutil/
	"os" //https://golang.org/pkg/os/
	"bufio"
)

func getSliceOfFileNames(dir string) []string {
        var sliceOfFileInfos []os.FileInfo
        var err error
	var fileNames []string

        sliceOfFileInfos, err = ioutil.ReadDir(dir)
        if (err != nil) {
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
		if err != nil { panic(err) }

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

type LinkInfo struct {
	RecipeName []string
	FileName []string
}

func (self LinkInfo) GenerateLinks() []string {
	var linkSlice []string

	for index, rName := range self.RecipeName {
		link := "<p><a href=\"" + self.FileName[index] + "\">" + rName + "</a></p>"

		linkSlice = append(linkSlice, link)
	}

	return linkSlice
}

func main() {
	
	var test = `{{.GenerateLinks}}`

	fNames := getSliceOfFileNames("../recipes/")

	rNames := getSliceOfRecipeNames("../recipes/")

	var linkInfo = LinkInfo{rNames, fNames}

	t := template.Must(template.New("links").Parse(test))

///	f.Funcs(template.FuncMap{"GenerateLinks"})

	err := t.Execute(os.Stdout, linkInfo)
	if err != nil { panic(err) }

/*
	fmt.Printf("\n\nRECIPE NAMES:\n\n")
	fmt.Println(linkInfo.RecipeName)

	fmt.Printf("\n\nFILE NAMES:\n\n")
	fmt.Println(linkInfo.FileName)

	fmt.Printf("\n\nLINK GENERATION:\n\n")
	fmt.Println(linkInfo.GenerateLinks())*/
}
