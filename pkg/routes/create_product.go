package routes

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/database"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func createProduct(ctx *gin.Context) {
	product := ordine.Product{}
	db, err := database.GetConnector()
	if err != nil {
		fmt.Printf("Error trying to connect to database %s", err.Error())
	}

	defer ctx.Request.Body.Close()
	err = utils.UnmarshalBody(ctx.Request.Body, &product)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error trying unmarshal request body")
	}

	db.Create(&product)

	ctx.JSON(http.StatusCreated, gin.H{
		"name": product.Name,
	})
}
