package rate

import (
	"crud/http/handlers"
	"crud/http/middleware"
	"crud/infrastructure/database"
	"crud/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRateRoutes(router *gin.Engine, db *sql.DB) {
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	rateUsecase := usecase.NewRateUsecase(
		database.NewRateDb(db),
		usecase.NewUserUsecase(database.NewUserDb(db)),
	)

	rateHandler := handlers.NewRateHandler(rateUsecase)

	protected.POST("/movie/rate", rateHandler.ActionRateMovie)
	protected.GET("/movie/rates/list", rateHandler.ActionListRates)
	protected.GET("/movie/rate/:id/details", rateHandler.ActionRateDetails)
}
