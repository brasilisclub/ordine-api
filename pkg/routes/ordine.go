package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ordine-api/pkg/ordine"

	"github.com/gin-gonic/gin"
)

func OrdineRoutes(r *gin.Engine) {
	r.GET("/ordine", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			gin.H{
				"name": "test",
			})
	})
	r.POST("/ordine", func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error reading body request")
		}

		ord := ordine.Ordine{}
		err = json.Unmarshal(body, &ord)
		if err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("Error trying unmarshal request body: %s", err.Error()))
		}

		ctx.JSON(http.StatusOK, gin.H{
			"table": ord.Table,
		})
	})
	r.PUT("/ordine/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.Param("id"))
	})
	r.DELETE("/ordine/:id", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.Param("id"))
	})
}
