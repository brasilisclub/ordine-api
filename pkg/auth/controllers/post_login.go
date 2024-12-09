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
//	@Summary		Login
//	@Description	Login for a User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		auth.LoginRequestBody	true	"User data"
//	@Success		200		{object}	auth.LoginUserSuccessResponse
//	@Failure		400		{object}	auth.AuthFailResponse
//	@Failure		500		{object}	auth.AuthFailResponse
//	@Router			/auth/login [post]
func PostLogin(ctx *gin.Context) {
	var requestBody auth.LoginRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, auth.AuthFailResponse{Message: fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	token, err := services.Login(&auth.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, auth.AuthFailResponse{
			Message: fmt.Sprintf("Error trying to generate token: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, auth.LoginUserSuccessResponse{
		Token: token,
	})
}
