package router

import (
	"backend/src/controllers"
	"backend/src/middlewares"
	"backend/src/repositories"
	"backend/src/services"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	// 環境変数からCORS設定を読み込む
	allowOrigins := os.Getenv("CORS_ALLOW_ORIGINS")
	origins := strings.Split(allowOrigins, ",")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/login", loginController.Login)
	r.POST("/answers", answersController.SaveAnswers)

	quizDataRouter := r.Group("/questions")
	quizDataRouter.GET("oneday", quizDataController.GetOneDaysQuiz) // 1日分のクイズデータ返却メソッド

	adminDataRouter := r.Group("/admins")
	adminDataRouter.Use(middlewares.AuthMiddleware()) // ミドルウェアを適用

	questionsDataRouter := adminDataRouter.Group("/questionsdata")
	questionsDataRouter.POST("/import_csv", adminsController.ImportCSV)  // 新規クイズデータ挿入メソッド
	questionsDataRouter.GET("/all", adminsController.FindAllQuestions)   // 全クイズデータ返却メソッド
	questionsDataRouter.PUT("/:id", adminsController.UpdateQuestions)    // クイズデータの更新
	questionsDataRouter.DELETE("/:id", adminsController.DeleteQuestions) // クイズデータの削除

	userDataRouter := adminDataRouter.Group("/userdata")
	userDataRouter.GET("/userslist", adminsController.GetUsersInfomation)  // ユーザーデータ一覧取得
	userDataRouter.PUT("/updateusers/:id", adminsController.UpdateUsers)   // ユーザーデータの更新
	userDataRouter.POST("/addusers", adminsController.AddUsers)            // ユーザーデータの追加
	userDataRouter.DELETE("/deleteuser/:id", adminsController.DeleteUsers) // ユーザーデータの削除

	analizeDataRouter := adminDataRouter.Group("/analizedata")
	analizeDataRouter.GET("/ranking", adminsController.GetRanking) // ランキングデータ一覧取得

	return r

}
