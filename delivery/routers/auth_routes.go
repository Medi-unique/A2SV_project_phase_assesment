package routers

import (
	"github.com/gin-gonic/gin"

	"assessment1/config/db"
	"assessment1/delivery/controllers"
	"assessment1/repository"
	"assessment1/usecase"
)

func SetUpAuth(router *gin.Engine) {

	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authController := controllers.NewUserController(userUsecase)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.GET("/activate", authController.ActivateAccount)

		// OAuth``

		// reset password
		auth.POST("/reset-password", authController.SendPasswordResetLink)
		auth.POST("/reset-password/:token", authController.ResetPassword)
	}
}
