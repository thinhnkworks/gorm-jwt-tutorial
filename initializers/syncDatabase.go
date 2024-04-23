package initializers

import "gin-jwt-tutorial/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
