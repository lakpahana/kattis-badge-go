package repo

import (
	"gorm.io/gorm"
	"lakpahana.me/db"
	"lakpahana.me/db/models"
)

func CreateOrUpdateUser(DB *gorm.DB, user models.User) string {
	existingUser := models.User{}
	result := DB.Where("username = ?", user.Username).First(&existingUser)

	if result.Error != nil {
		DB.Create(&user)
		return "OK"
	} else {
		// If user found, replace with provided details
		user.ID = existingUser.ID // Retain the existing user's ID
		DB.Save(&user)
		return "OK"
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
