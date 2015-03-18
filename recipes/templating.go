package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name     string
	FavColor string
}

func main() {
	const message = `
Hello there, {{.Name}}!
I'm told your favorite color is {{.FavColor}}.
{{if eq .FavColor "Blue"}} Mine too!!! {{else}}That's nice I guess.{{end}}

`
	var p = Person{"Corey", "Red"}

	var t = template.New("greeting")
	t.Parse(message)
	err := t.Execute(os.Stdout, p)

	if err != nil {
		fmt.Printf("ERROR:%v", err)
	}
}
