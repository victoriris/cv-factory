package main

import (
	"context"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/victormidp/cv-factory/prisma/db"
)

func getUsers(c *gin.Context) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	ctx := context.Background()

	users, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, users)
}

func createDocument(c *gin.Context) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	ctx := context.Background()

	user, err := client.User.FindFirst(
		db.User.ID.Equals("ckrii17890000m6vr4idf6yxj"),
	).With(
		db.User.Jobs.Fetch(),
		db.User.Skills.Fetch(),
		db.User.List.Fetch(),
		db.User.EducationItems.Fetch(),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("example/index.html"))
	cv, err := os.Create("example/output.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(cv, user)
	c.IndentedJSON(http.StatusOK, user)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	// Document
	router.POST("/document", createDocument)

	// router.POST("/job", todo)
	// router.POST("/education-item", todo)
	// router.POST("/skill", todo)
	// router.POST("/list", todo)
	// router.POST("/login", todo)
	// router.POST("/refresh", todo)

	// Run service
	router.Run(("localhost:8000"))
}
