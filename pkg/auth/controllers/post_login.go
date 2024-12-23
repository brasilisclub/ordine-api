package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/auth/services"
	"ordine-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// PostLogin godoc
//
//	@Summary		Login
//	@Description	Login for a User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		auth.AuthRequestBody	true	"User data"
//	@Success		200		{object}	auth.LoginUserSuccessResponse
//	@Failure		400		{object}	utils.GenericResponse
//	@Failure		500		{object}	utils.GenericResponse
//	@Router			/auth/login [post]
func PostLogin(ctx *gin.Context) {
	var requestBody auth.AuthRequestBody

	if err := ctx.ShouldBindBodyWithJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.GenericResponse{Message: fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}

	token, err := services.Login(&auth.User{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		if errors.Is(err, auth.ErrorUserNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			ctx.JSON(http.StatusUnauthorized, utils.GenericResponse{
				Message: fmt.Sprintf("Invalid credentials: %s", err.Error()),
			})
		}

		ctx.JSON(http.StatusInternalServerError, utils.GenericResponse{
			Message: fmt.Sprintf("Error trying to generate token: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, auth.LoginUserSuccessResponse{
		Token: token,
	})
}
