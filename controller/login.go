package controller

import (
	"capstone.com/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type login struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var login login
	var user model.User
	db.Where("email = ?", login.Email).Where("password = ?", login.Password).Find(&user)
	if user.Email != login.Email {
		c.JSON(400, gin.H{
			"status": "email atau password salah",
		})
	}
	if user.Level == 1 {
		c.JSON(200, gin.H{
			"status": "berhasil",
			"level":  "admin",
			"email":  user.Email,
		})
	}
	if user.Level == 2 {
		c.JSON(200, gin.H{
			"status": "berhasil",
			"level":  "dokter",
			"email":  user.Email,
		})
	}
	if user.Level == 3 {
		c.JSON(200, gin.H{
			"status": "berhasil",
			"level":  "perawat",
			"email":  user.Email,
		})
	}
}
