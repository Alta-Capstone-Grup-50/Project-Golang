package main

import (
	"fmt"
	"os"

	"capstone.com/controller"
	"capstone.com/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	godotenv.Load()
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	fmt.Println(address)

	r.Run(address)
}
