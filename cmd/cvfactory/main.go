package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

type EducationItem struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	InstitutionName string `json:"institutionName"`
	Location        string `json:"location"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	ReferenceId     string `json:"referenceId"`
	Url             string `json:"url"`
}

type List struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Skill struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Proficiency string `json:"proficiency"`
}

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CompanyName string `json:"companyName"`
	Location    string `json:"location"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type UserData struct {
	Firstname      string          `json:"firstname"`
	Lastname       string          `json:"lastname"`
	Email          string          `json:"email"`
	Phone          string          `json:"phone"`
	Address        string          `json:"address"`
	Jobs           []Job           `json:"jobs"`
	Skills         []Skill         `json:"skills"`
	Lists          []List          `json:"lists"`
	EducationItems []EducationItem `json:"educationItems"`
}

func FillTemplate(dataPath string, templatePath string) {
	file, _ := ioutil.ReadFile(dataPath)
	data := UserData{}
	_ = json.Unmarshal([]byte(file), &data)

	tmpl := template.Must(template.ParseFiles(templatePath))
	outputPath := filepath.Dir(templatePath) + "/output.html"
	cv, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(cv, data)
	fmt.Println("Templated filled!✨")
	fmt.Println(outputPath)
}

func main() {
	argsWithoutProg := os.Args[1:]
	command := argsWithoutProg[0]

	if command == "fill" {
		// "sample-data.json"
		dataPath := argsWithoutProg[1]

		// "example/index.html"
		templatePath := argsWithoutProg[2]
		FillTemplate(dataPath, templatePath)
	}
}
