package routes

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createProduct(ctx *gin.Context) {
	product := ordine.Product{}
	dsn := "root:root@tcp(database:3306)/ordine?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error trying to connect to database %s", err.Error())
	}
	db.AutoMigrate(&ordine.Product{})

	defer ctx.Request.Body.Close()
	err = utils.UnmarshalBody(ctx.Request.Body, &product)
	db.Create(&product)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Error trying unmarshal request body")
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"name": product.Name,
	})
}
