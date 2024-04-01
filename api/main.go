package main

import (
	"github.com/gin-gonic/gin"
	"lakpahana.me/api/controllers"
	"lakpahana.me/db"
)

func init() {
	db.ConnectToDB()
}

func main() {

	router := gin.Default()
	
	router.GET("/allUsers", controllers.GetAllUsers)
	router.GET("/rank/:id", controllers.GetRankByCountry)
	// router.GET("")
	router.Run(":8080")

}
