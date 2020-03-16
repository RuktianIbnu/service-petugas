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

	router.POST("/service-petugas/simpanperubahan-data", inDB.InsertDataPerubahan)
	router.Run(":3400")
}