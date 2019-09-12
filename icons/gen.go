// The following directive necessary to make the package coherent:

// +build ignore

// This program generates map.go. You can invoke it by running
// go generate

package icons

import (
	"log"
	"os"
	"text/template"
	"time"
)

type MapTemplate struct {
	timestamp time.Time
	source    string
	data      map[string]string
}

func main() {
	tmpl, err := template.ParseFiles("map.go.tmpl")
	if err != nil {
		log.Fatalf("An error occured when parsing template: %v", err)
	}

	file, err := os.Create("map.go")
	if err != nil {
		log.Fatalf("Error with creating file: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file)
}
