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

	// クイズデータ取得用
	quizDataRepository := repositories.NewQuestionsRepository(db)
	quizDataService := services.NewQuestionsService(quizDataRepository)
	quizDataController := controllers.NewQuestionsController(quizDataService)

	// ログイン用
	loginRepository := repositories.NewLoginRepository(db)
	loginService := services.NewLoginService(loginRepository)
	loginController := controllers.NewLoginController(loginService)

	r := gin.Default()

	r.Use(cors.Default()) //フロントとやり取りする場合は設定したほうがよい。

	r.POST("/login", loginController.Login)

	quizDataRouter := r.Group("/questions")

	quizDataRouter.GET("", quizDataController.FindAll)              // 全クイズデータ返却メソッド
	quizDataRouter.GET("oneday", quizDataController.GetOneDaysQuiz) // 1日分のクイズデータ返却メソッド
	quizDataRouter.GET("/:id", quizDataController.FindById)
	quizDataRouter.POST("", quizDataController.Create)
	quizDataRouter.PUT("/:id", quizDataController.Update)
	quizDataRouter.DELETE("/:id", quizDataController.Delete)

	return r

}
