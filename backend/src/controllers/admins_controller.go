package controllers

import (
	"backend/src/dto"
	"backend/src/services"
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
	AddUsers(ctx *gin.Context)
	DeleteUsers(ctx *gin.Context)
	GetRanking(ctx *gin.Context)     // ランキング取得用　後で消す
	GetGraphData(ctx *gin.Context)   // 個人の成績グラフ取得用
	GetInitialData(ctx *gin.Context) // ランキングと全体の傾向グラフ取得用
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dbId"})
		return
	}

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

func (c *AdminsController) AddUsers(ctx *gin.Context) {
	var addUsers dto.AdmUserData
	if err := ctx.ShouldBindJSON(&addUsers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addedUsers, err := c.service.AddUsers(addUsers)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addedUsers)
}

func (c *AdminsController) DeleteUsers(ctx *gin.Context) {
	dbId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dbId"})
		return
	}

	err = c.service.DeleteUsers(uint(dbId))
	if err != nil {
		if err.Error() == "User not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.Status(http.StatusOK)
}

// 後でなしにする
func (c *AdminsController) GetRanking(ctx *gin.Context) {
	ranking, err := c.service.GetRanking()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ranking": ranking,
	})
}

// ユーザー成績のグラフデータを返す
func (c *AdminsController) GetGraphData(ctx *gin.Context) {
	Id := ctx.Param("id")

	graphData, err := c.service.GetGraphData(Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get graph data"})
		return
	}

	ctx.JSON(http.StatusOK, graphData)
}

// ランキングと全体の傾向グラフデータを返す
func (c *AdminsController) GetInitialData(ctx *gin.Context) {
	Id := ctx.Param("id")

	// ランキングのデータを取得
	ranking, err := c.service.GetRanking()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	// 全体の傾向グラフのデータを取得
	graphData, err := c.service.GetGraphData(Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get graph data"})
		return
	}

	// JSONで返す
	ctx.JSON(http.StatusOK, gin.H{
		"ranking":   ranking,
		"graphData": graphData,
	})
}
