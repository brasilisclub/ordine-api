package routes

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/ordine"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func createOrdine(ctx *gin.Context) {
	ord := ordine.Ordine{}

	defer ctx.Request.Body.Close()
	err := utils.UnmarshalBody(ctx.Request.Body, &ord)

	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error trying unmarshal request body: %s", err.Error()))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"table": ord.Table,
	})
}
