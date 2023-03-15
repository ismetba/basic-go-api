package initializers

import "github.com/ismetbayandur/goapi/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}
