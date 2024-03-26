package repo

import (
	"gorm.io/gorm"
	"lakpahana.me/db"
	"lakpahana.me/db/models"
)

func CreateOrUpdateUser(DB *gorm.DB, user models.User) string {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return "OK"
	} else {
		return "Error"
	}
}

func GetUserById(username string) models.User {
	var result models.User
	// fmt.Println(username)
	db.DB.Find(&result, "username = ?", username)
	// fmt.Print(result)
	return result
}

func GetAllUsers() []models.User {
	var result []models.User
	db.DB.Find(&result)
	return result
}
