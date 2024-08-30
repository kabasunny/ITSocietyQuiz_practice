package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAnswersController interface {
	SaveAnswers(ctx *gin.Context)
}

type AnswersController struct {
	service services.IAnswersService
}

func NewAnswersController(service services.IAnswersService) IAnswersController {
	return &AnswersController{service: service}
}

func (c *AnswersController) SaveAnswers(ctx *gin.Context) {
	var input dto.AnswersInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// トークンの確認
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
		return
	}

	// 回答の保存
	err := c.service.SaveAnswers(input, token)
	if err != nil {
		if err.Error() == "invalid token" || err.Error() == "invalid empID" || err.Error() == "token has expired" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Answers saved successfully"})
}
