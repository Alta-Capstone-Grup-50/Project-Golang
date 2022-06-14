package controller

import (
	"capstone.com/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type post_pasien struct {
	Nama           string `form:"nama"`
	Nik            string `form:"nik"`
	Alamat         string `form:"alamat"`
	Jenis_kelamin  string `form:"jenis_kelamin"`
	No_hp          string `form:"no_hp"`
	Tempat_lahir   string `form:"tempat_lahir"`
	Tanggal_lahir  string `form:"tanggal_lahir"`
	Id_rekam_medis int    `form:"id_rekam_medis"`
}

type get_pasien struct {
	Id            int    `form:"id"`
	Nama          string `form:"nama"`
	Nik           string `form:"nik"`
	Alamat        string `form:"alamat"`
	Jenis_kelamin string `form:"jenis_kelamin"`
	Rekam_medis   interface{}
}

type rekam_medis struct {
	Pemeriksaan      string `form:"pemeriksaan"`
	Jenis_penanganan string `form:"jenis_penanganan"`
}

func Get_pasien(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var get_pasien []get_pasien
	db.Raw("SELECT id, nama, nik, alamat, jenis_kelamin FROM capstone.pasiens;").Scan(&get_pasien)
	for i := 0; i < len(get_pasien); i++ {
		var rekam_medis rekam_medis
		db.Where("id_pasien = ?", get_pasien[i].Id).Find(&rekam_medis)
		db.Model(&get_pasien[i]).Update("Rekam_medis", rekam_medis)
	}
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
		Nik:            post.Nik,
		Alamat:         post.Alamat,
		Jenis_kelamin:  post.Jenis_kelamin,
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
