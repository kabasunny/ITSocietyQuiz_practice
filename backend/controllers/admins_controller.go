package controllers

import (
	"backend/dto"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IAdminsController interface {
	FindAllQuestions(ctx *gin.Context)
	FindQuestionsById(ctx *gin.Context)
	// CreateQuestions(ctx *gin.Context)
	UpdateQuestions(ctx *gin.Context)
	DeleteQuestions(ctx *gin.Context)
	ImportCSV(ctx *gin.Context)
	GetUsersInfomation(ctx *gin.Context)
	UpdateUsers(ctx *gin.Context)
}

type AdminsController struct {
	service services.IAdminsService
}

func NewAdminsController(service services.IAdminsService) IAdminsController {
	return &AdminsController{service: service}
}

func (c *AdminsController) FindAllQuestions(ctx *gin.Context) {
	Questions, err := c.service.FindAllQuestions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"adm_data": Questions})
}

func (c *AdminsController) FindQuestionsById(ctx *gin.Context) {
	QuestionsId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	Questions, err := c.service.FindQuestionsById(uint(QuestionsId))
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

// func (c *AdminsController) CreateQuestions(ctx *gin.Context) {
// 	var input dto.CreateQuestionsInput
// 	if err := ctx.ShouldBindJSON(&input); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	newQuestions, err := c.service.Create(input)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, gin.H{"data": newQuestions})
// }

func (c *AdminsController) UpdateQuestions(ctx *gin.Context) {
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

	// IDの確認
	if uint(QuestionsId) != input.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID mismatch"})
		return
	}

	updateQuestions, err := c.service.UpdateQuestions(uint(QuestionsId), input)
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

func (c *AdminsController) DeleteQuestions(ctx *gin.Context) {
	QuestionsId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = c.service.DeleteQuestions(uint(QuestionsId))
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

func (c *AdminsController) ImportCSV(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	filePath := "./" + file.Filename
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// データ処理をサービス層に委譲
	if err := c.service.ProcessCSVData(filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded and data processed successfully"})
}

func (c *AdminsController) GetUsersInfomation(ctx *gin.Context) {
	userList, err := c.service.GetUsersInfomation()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	// todaysfinish = false // テスト時の制限解除はサービス側で行う

	ctx.JSON(http.StatusOK, gin.H{
		"userlist": userList, // レスポンスにtodays_finishを含める
	})
}

func (c *AdminsController) UpdateUsers(ctx *gin.Context) {
	dbId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var updateUsers dto.AdmUserData
	if err := ctx.ShouldBindJSON(&updateUsers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := c.service.UpdateUsers(uint(dbId), updateUsers)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
