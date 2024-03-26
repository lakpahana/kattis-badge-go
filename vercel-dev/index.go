package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lakpahana.me/api/controllers"
	"lakpahana.me/db"
)

// func init() {
// }

// func main() {

// }

func Handler(w http.ResponseWriter, r *http.Request) {
	// app.ServeHTTP(w,r)
	db.ConnectToDB()
	router := gin.Default()
	router.GET("/allUsers", controllers.GetAllUsers)
	router.GET("/rank/:id", controllers.GetRankByCountry)
	// router.GET("")
	router.Run("localhost:8080")
}
