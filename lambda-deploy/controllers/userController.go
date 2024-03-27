package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lakpahana.me/db/repo"
)

func Hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "wroking"})
}

func GetAllUsers(c *gin.Context) {
	// var users []models.User
	result := repo.GetAllUsers()
	c.IndentedJSON(http.StatusOK, gin.H{"message": result})
}

func GetRankByCountry(c *gin.Context) {
	id := c.Param("id")

	result := repo.GetUserById(id)
	c.IndentedJSON(http.StatusOK, gin.H{"message": result.Country_rank})

}
