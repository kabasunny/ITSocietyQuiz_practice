package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IQuestionsController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetOneDaysQuiz(ctx *gin.Context)
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

func (c *QuestionsController) Create(ctx *gin.Context) {
	var input dto.CreateQuestionsInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newQuestions, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newQuestions})
}

func (c *QuestionsController) Update(ctx *gin.Context) {
	QuestionsId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var input dto.UpdateQuestionsInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateQuestions, err := c.service.Update(uint(QuestionsId), input)
	if err != nil {
		if err.Error() == "Questions not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updateQuestions})
}

func (c *QuestionsController) Delete(ctx *gin.Context) {
	QuestionsId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = c.service.Delete(uint(QuestionsId))
	if err != nil {
		if err.Error() == "Questions not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.Status(http.StatusOK)

}

func (c *QuestionsController) GetOneDaysQuiz(ctx *gin.Context) {
	NunmberOfQuestions := uint(5) // 1日あたりの問題数を入力
	QuizDatas, err := c.service.GetOneDaysQuiz(NunmberOfQuestions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": QuizDatas})
}
