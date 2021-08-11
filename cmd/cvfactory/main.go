package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
)

type EducationItem struct {
	Title           string `json: "title"`
	Description     string `json: "description"`
	InstitutionName string `json: "institutionName"`
	Location        string `json: "location"`
	StartDate       string `json: "startDate"`
	EndDate         string `json: "endDate"`
	ReferenceId     string `json: "referenceId"`
	Url             string `json: "url"`
}

type List struct {
	Name    string `json: "name"`
	Content string `json: "content"`
}

type Skill struct {
	Name        string `json: "name"`
	Description string `json: "description"`
	Proficiency string `json: "proficiency"`
}

type Job struct {
	Title       string `json: "title"`
	Description string `json: "description"`
	CompanyName string `json: "companyName"`
	Location    string `json: "location"`
	StartDate   string `json: "startDate"`
	EndDate     string `json: "endDate"`
}

type UserData struct {
	Firstname      string          `json: "firstname"`
	Lastname       string          `json: "lastname"`
	Email          string          `json: "email"`
	Phone          string          `json: "phone"`
	Address        string          `json: "address"`
	Jobs           []Job           `json: "jobs"`
	Skills         []Skill         `json: "skills"`
	Lists          []List          `json: "lists"`
	EducationItems []EducationItem `json: "educationItems"`
}

func main() {
	file, _ := ioutil.ReadFile("sample-data.json")
	data := UserData{}
	_ = json.Unmarshal([]byte(file), &data)

	tmpl := template.Must(template.ParseFiles("example/index.html"))
	cv, err := os.Create("example/output.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(cv, data)
}
