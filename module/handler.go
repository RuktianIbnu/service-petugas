package module

import (
	"fmt"
	"../structs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
)

func CekAuth(c *gin.Context) bool{
	var status bool
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		status = true
	} else {
		status = false
	}
	return status
}

func (idb *InDB) ExeInsertSpesifikasi(dataSpesifikasi structs.Spesifikasi, c *gin.Context) string{
    success := idb.DB.Create(&dataSpesifikasi).Error
	hasil 	:= "success"

		if success != nil {
			hasil = "error"
		}
	return hasil
}

func (idb *InDB) ExeInsertManagement(dataManagement structs.Management, c *gin.Context) string{
    success := idb.DB.Create(&dataManagement).Error
	hasil 	:= "success"

		if success != nil {
			hasil = "error"
		}
	return hasil
}

func (idb *InDB) ExeInsertIP(dataIP structs.Ip_Address, c *gin.Context) string{
    success := idb.DB.Create(&dataIP).Error
	hasil 	:= "success"

		if success != nil {
			hasil = "error"
		}
	return hasil
}

func (idb *InDB) ExeInsertReview(data structs.ReviewPetugas, c *gin.Context) string{
	success := idb.DB.Create(&data).Error
	hasil 	:= "success"

        if success != nil {
            hasil = "error"
		}
	return hasil
}