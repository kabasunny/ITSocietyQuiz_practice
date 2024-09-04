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

	// 解答格納用
	answersRepository := repositories.NewAnswersRepository(db)
	answersService := services.NewAnswersService(answersRepository)
	answersController := controllers.NewAnswersController(answersService)

	r := gin.Default()

	// カスタムCORS設定を追加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3005"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/login", loginController.Login)

	r.POST("/answers", answersController.SaveAnswers)

	quizDataRouter := r.Group("/questions")

	quizDataRouter.GET("", quizDataController.FindAll)              // 全クイズデータ返却メソッド
	quizDataRouter.GET("oneday", quizDataController.GetOneDaysQuiz) // 1日分のクイズデータ返却メソッド
	quizDataRouter.GET("/:id", quizDataController.FindById)
	quizDataRouter.POST("", quizDataController.Create)
	quizDataRouter.PUT("/:id", quizDataController.Update)
	quizDataRouter.DELETE("/:id", quizDataController.Delete)

	return r
}
