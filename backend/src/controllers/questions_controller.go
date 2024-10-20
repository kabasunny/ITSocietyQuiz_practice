package controllers

import (
	"backend/src/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IQuestionsController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	GetOneDaysQuiz(ctx *gin.Context) // 1日分のクイズを取得する
}

type QuestionsController struct {
	service services.IQuestionsService
}

func NewQuestionsController(service services.IQuestionsService) IQuestionsController {
	return &QuestionsController{service: service}
}

func (c *QuestionsController) FindAll(ctx *gin.Context) {
	Questions, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": Questions})
}

func (c *QuestionsController) FindById(ctx *gin.Context) {
	QuestionsId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	Questions, err := c.service.FindById(uint(QuestionsId))
	if err != nil {
		if err.Error() == "Questions not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": Questions})
}

func (c *QuestionsController) GetOneDaysQuiz(ctx *gin.Context) {
	// トークンの確認
	tokenString := ctx.GetHeader("Authorization")

	// クエリパラメータからtodays_countを取得
	todaysCountStr := ctx.Query("todays_count")

	// todays_count を uint に変換
	todaysCountInt, err := strconv.Atoi(todaysCountStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todays_count"})
		return
	}

	todaysCount := uint(todaysCountInt)

	QuizDatas, todaysfinish, err := c.service.GetOneDaysQuiz(tokenString, todaysCount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	// todaysfinish = false // テスト時の制限解除はサービス側で行う

	ctx.JSON(http.StatusOK, gin.H{
		"quizdata":      QuizDatas,
		"todays_finish": todaysfinish, // レスポンスにtodays_finishを含める
	})
}
