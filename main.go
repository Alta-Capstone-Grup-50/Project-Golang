package main

import (
	"capstone.com/controller"
	"capstone.com/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := model.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/pasien", controller.Get_pasien)
	r.POST("/pasien", controller.Post_pasien)
	r.GET("/login", controller.Login)
	r.Run()
}
