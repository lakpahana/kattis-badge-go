package main

import (
	"lakpahana.me/db"
	"lakpahana.me/db/models"
)


func main() {

	db.ConnectToDB()
	db.DB.AutoMigrate(&models.User{})
}
