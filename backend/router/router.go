package router

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/repositories"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	// ユーザーのクイズデータ取得用
	quizDataRepository := repositories.NewQuestionsRepository(db)
	quizDataService := services.NewQuestionsService(quizDataRepository)
	quizDataController := controllers.NewQuestionsController(quizDataService)

	// ユーザーの解答格納用
	answersRepository := repositories.NewAnswersRepository(db)
	answersService := services.NewAnswersService(answersRepository)
	answersController := controllers.NewAnswersController(answersService)

	// ユーザーおよび管理者のログイン用
	loginRepository := repositories.NewLoginRepository(db)
	loginService := services.NewLoginService(loginRepository)
	loginController := controllers.NewLoginController(loginService)

	// 管理者のデータ編集用
	adminsRepository := repositories.NewAdminsRepository(db)
	adminsService := services.NewAdminsService(adminsRepository)
	adminsController := controllers.NewAdminsController(adminsService)

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

	adminDataRouter := r.Group("/admins")
	adminDataRouter.Use(middlewares.AuthMiddleware()) // ミドルウェアを適用

	// quizDataRouter.GET("", quizDataController.FindAll)              // 全クイズデータ返却メソッド
	quizDataRouter.GET("oneday", quizDataController.GetOneDaysQuiz) // 1日分のクイズデータ返却メソッド
	// quizDataRouter.GET("/:id", quizDataController.FindById)

	adminDataRouter.POST("/import_csv", adminsController.ImportCSV)  // 新規クイズデータ挿入メソッド
	adminDataRouter.GET("", adminsController.FindAllQuestions)       // 全クイズデータ返却メソッド
	adminDataRouter.PUT("/:id", adminsController.UpdateQuestions)    // クイズデータの更新
	adminDataRouter.DELETE("/:id", adminsController.DeleteQuestions) // クイズデータの削除

	adminDataRouter.GET("/userslist", adminsController.GetUsersInfomation) // ユーザーデータ一覧取得
	adminDataRouter.PUT("/updateusers/:id", adminsController.UpdateUsers)  // ユーザーデータの更新

	return r
}
