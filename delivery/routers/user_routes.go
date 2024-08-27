package routers

import (
	"github.com/gin-gonic/gin"

	"assessment1/config/db"
	"assessment1/delivery/controllers"
	"assessment1/infrastracture"
	"assessment1/repository"
	"assessment1/usecase"
)

func SetUpUser(router *gin.Engine) {
	//user routes
	userRepo := repository.NewUserRepositoryImpl(db.UserCollection)
	userUsecase := usecase.NewUserUsecase(userRepo)
	authController := controllers.NewUserController(userUsecase)
	loanRepo := repository.NewLoanRepositoryImpl(db.LoanCollection)
	loanUsecase := usecase.NewLoanUsecase(loanRepo)
	loanController := controllers.NewLoanController(loanUsecase)
	user := router.Group("/user")
	user.Use(infrastracture.AuthMiddleware())

	{
		user.GET("/me", authController.GetMyProfile)
		user.PUT("/update", authController.UpdateMyProfile)
		user.POST("/upload-image", authController.UploadImage)
		user.DELETE("/me", authController.DeleteMyAccount)

		// Logout Routes

		user.POST("/refresh-token", authController.RefreshToken)

		user.POST("/logout", authController.Logout)
		user.GET("logout-all", authController.LogoutAll)
		// Loan Routes
		user.POST("/loans", loanController.ApplyForLoan)
		user.GET("/loans/:id", loanController.GetLoanStatus)
	}
}
