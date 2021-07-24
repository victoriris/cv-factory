package main

import (
	"context"
	"net/http"

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

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	// router.POST("/create-document", todo)
	// router.POST("/job", todo)
	// router.POST("/education-item", todo)
	// router.POST("/skill", todo)
	// router.POST("/list", todo)
	// router.POST("/login", todo)
	// router.POST("/refresh", todo)

	// Run service
	router.Run(("localhost:8000"))
}

// TODO: handle template writing
// Open JSON file
// jsonData, err := ioutil.ReadFile("data.json")
// if err != nil {
// 	panic(err)
// }

// var page PageData
// tmpl := template.Must(template.ParseFiles("template.html"))
// json.Unmarshal([]byte(jsonData), &page)

// cv, err := os.Create("output/cv.html")
// if err != nil {
// 	panic(err)
// }
// tmpl.Execute(cv, jsonData)
