package main

import (
	"./config"
	"./module"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DBInit()
	inDB := &module.InDB{DB: db}
	router := gin.Default()

	router.POST("/service-petugas/simpan-DataPerubahan", inDB.InsertDataPerubahan)
	router.POST("/service-petugas/simpan-DataReview", inDB.InsertReview2)
	router.GET("/service-petugas/Cari-SemuaDataReview", inDB.GetAllData)
	router.Run(":3400")
}