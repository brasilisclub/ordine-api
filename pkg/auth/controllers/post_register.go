package controllers

import (
	"fmt"
	"net/http"
	"ordine-api/pkg/auth"
	"ordine-api/pkg/auth/services"

	"github.com/gin-gonic/gin"
)

func PostRegister(ctx *gin.Context) {
	var bodyUser auth.User

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
