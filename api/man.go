package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID               int     `json:"id"`
	Username         string  `json:"username"`
	Name             string  `json:"name"`
	Rank             int     `json:"rank"`
	Score            float64 `json:"score"`
	Country          string  `json:"country"`
	Country_code     string  `json:"country_code"`
	Country_rank     int     `json:"country_rank"`
	Subdivision      string  `json:"subdiv"`
	Subdivision_code string  `json:"subdiv_code"`
	Subdivision_rank int     `json:"subdiv_rank"`
	University       string  `json:"uni"`
	University_code  string  `json:"uni_code"`
	University_rank  int     `json:"uni_rank"`
}

var users = []user{
	{
		ID:               1,
		Username:         "pramitha-jayasooriya",
		Name:             "Pramitha Jayasooriya",
		Rank:             5118,
		Score:            211.0,
		Country:          "Sri Lanka",
		Country_code:     "LKA",
		Country_rank:     1,
		Subdivision:      "",
		Subdivision_code: "",
		Subdivision_rank: 0,
		University:       "University of Ruhuna",
		University_code:  "ruh.ac.lk",
		University_rank:  1,
	},
	{
		ID:               2,
		Username:         "pramitha-jayasooriya",
		Name:             "Pramitha Jayasooriya",
		Rank:             5118,
		Score:            211.0,
		Country:          "Sri Lanka",
		Country_code:     "LKA",
		Country_rank:     1,
		Subdivision:      "",
		Subdivision_code: "",
		Subdivision_rank: 0,
		University:       "University of Ruhuna",
		University_code:  "ruh.ac.lk",
		University_rank:  1,
	},
}

func userById(c *gin.Context) {
	id := c.Param("id")
	book, err := getuserById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getuserById(id string) (*user, error) {
	for i, b := range users {
		if b.Username == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "yesss")
}

func main() {
	router := gin.Default()
	router.GET("user/:id", userById)
	router.GET("/", hello)
	router.Run("localhost:8080")
}
