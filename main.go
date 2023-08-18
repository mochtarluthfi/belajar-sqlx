package main

import (
	"belajar-sqlx/controllers"
	"belajar-sqlx/db_client"
	"github.com/gin-gonic/gin"
)

func main() {

	db_client.InitialiseDBConnection()
	r := gin.Default()

	r.POST("/", controllers.GetAllUser)
	r.POST("/get-one", controllers.CreatePost)

	if err := r.Run(":5000"); err != nil {
		panic(err.Error())
	}
}
