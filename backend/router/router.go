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

	quizDataRepository := repositories.NewQuizDataRepository(db)
	quizDataService := services.NewQuizDataService(quizDataRepository)
	quizDataController := controllers.NewQuizDataController(quizDataService)

	r := gin.Default()

	r.Use(cors.Default()) //フロントとやり取りする場合は設定したほうがよい。

	quizDataRouter := r.Group("/quiz_data")

	quizDataRouter.GET("", quizDataController.FindAll)
	quizDataRouter.GET("/:id", quizDataController.FindById)
	quizDataRouter.POST("", quizDataController.Create)
	quizDataRouter.PUT("/:id", quizDataController.Update)
	quizDataRouter.DELETE("/:id", quizDataController.Delete)

	return r

}
