package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"lakpahana.me/api/controllers"
	"lakpahana.me/db"
)


func init() {
	db.ConnectToDB()
}

func main() {

	router := gin.Default()
	router.GET("user/:id", userById)
	router.GET("/", hello)
	router.GET("/allUsers", controllers.GetAllUsers)
	router.GET("/rank/:id", controllers.GetRankByCountry)
	router.Run("localhost:8080")

}
