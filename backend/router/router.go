package router

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// 初期テスト用
	// quizDataRepository := repositories.NewQuizDataRepository(db)
	// quizDataService := services.NewQuizDataService(quizDataRepository)
	// quizDataController := controllers.NewQuizDataController(quizDataService)

	quizDataRepository := repositories.NewQuestionsRepository(db)
	quizDataService := services.NewQuestionsService(quizDataRepository)
	quizDataController := controllers.NewQuestionsController(quizDataService)

	r := gin.Default()

	r.Use(cors.Default()) //フロントとやり取りする場合は設定したほうがよい。

	// 初期テスト用
	// quizDataRouter := r.Group("/quiz_data")
	quizDataRouter := r.Group("/questions")

	quizDataRouter.GET("", quizDataController.FindAll)              // 全クイズデータ返却メソッド
	quizDataRouter.GET("oneday", quizDataController.GetOneDaysQuiz) // 1日分のクイズデータ返却メソッド
	quizDataRouter.GET("/:id", quizDataController.FindById)
	quizDataRouter.POST("", quizDataController.Create)
	quizDataRouter.PUT("/:id", quizDataController.Update)
	quizDataRouter.DELETE("/:id", quizDataController.Delete)

	return r

}
