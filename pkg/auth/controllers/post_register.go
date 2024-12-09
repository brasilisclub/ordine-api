package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/auth/services"

	"github.com/gin-gonic/gin"
)

// PostLogin godoc
//
//	@Summary		Register
//	@Description	Register for a User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		auth.LoginRequestBody	true	"User data"
//	@Success		201		{object}	auth.User
//	@Failure		400		{object}	utils.GenericResponse
//	@Failure		500		{object}	utils.GenericResponse
//	@Router			/auth/register [post]
func PostRegister(ctx *gin.Context) {
	var bodyUser auth.LoginRequestBody

	err := ctx.ShouldBindBodyWithJSON(&bodyUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid body: %s", err.Error()),
		})
		return
	}
	err = services.CreateUser(&bodyUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Internal error trying to create user %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, bodyUser)
}
