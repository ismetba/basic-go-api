package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ismetbayandur/goapi/initializers"
	"github.com/ismetbayandur/goapi/controllers"
	)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.Run()
}
