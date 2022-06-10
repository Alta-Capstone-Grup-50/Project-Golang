package controller

import (
	"capstone.com/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type post_pasien struct {
	Nama           string `json:"nama"`
	No_identitas   string `json:"no_identitas"`
	Alamat         string `json:"alamat"`
	Gender         string `json:"gender"`
	No_hp          string `json:"no_hp"`
	Tempat_lahir   string `json:"tempat_lahir"`
	Tanggal_lahir  string `json:"tanggal_lahir"`
	Id_rekam_medis int    `json:"id_rekam_medis"`
}

func Get_pasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var get_pasien []model.Pasien
	db.Find(&get_pasien)
	c.JSON(200, gin.H{
		"data":   get_pasien,
		"status": "berhasil",
	})
}

func Post_pasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var post post_pasien
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "input tidak dalam bentuk json",
		})
		return
	}
	new := model.Pasien{
		Nama:           post.Nama,
		No_identitas:   post.No_identitas,
		Alamat:         post.Alamat,
		Gender:         post.Gender,
		No_hp:          post.No_hp,
		Tempat_lahir:   post.Tempat_lahir,
		Tanggal_lahir:  post.Tanggal_lahir,
		Id_rekam_medis: post.Id_rekam_medis,
	}
	db.Create(&new)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   new,
	})
}
