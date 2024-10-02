package controllers

import (
	"backend/src/dto"
	"backend/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	Login(ctx *gin.Context)
}

type LoginController struct {
	service services.ILoginService
}

func NewLoginController(service services.ILoginService) ILoginController {
	return &LoginController{service: service}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, isAdmin, todaysCount, err := c.service.Login(input.EmpID, input.Password)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if token == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token is nil"})
		return
	}

	response := dto.LoginResponse{
		Token:       *token,
		Admin:       isAdmin,
		TodaysCount: todaysCount, // 本日何問受けているか
	}
	ctx.JSON(http.StatusOK, response)
}
