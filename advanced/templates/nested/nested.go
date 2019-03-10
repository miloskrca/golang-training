package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseGlob("./demo/templates/nested/templates/*.tmpl"))
	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
