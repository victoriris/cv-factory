package main

import (
	"os"
	"path/filepath"

	server "github.com/victoridp/cv-factory/pkg/fileserver"
	template "github.com/victoridp/cv-factory/pkg/template"
)

func main() {
	argsWithoutProg := os.Args[1:]
	command := argsWithoutProg[0]

	if command == "fill" {
		// "sample-data.json"
		dataPath := argsWithoutProg[1]
		// "example/index.html"
		templatePath := argsWithoutProg[2]

		outputName := "output.html"
		outpath, err := template.FillTemplate(dataPath, templatePath, outputName)
		if err != nil {
			panic(err)
		}
		server.ServeFiles(filepath.Dir(outpath), outputName, 3000)
	}
}
