// Delivery/routers/router.go
package routers

import (
	"github.com/gin-gonic/gin"

)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// auth routes
	SetUpAuth(router)

	// // user routes
	SetUpUser(router)
	// // Admin routes
	// SetUpAdmin(router)

	// // oauth routes


	// // Admin routes
	SetUpAdmin(router)

	return router
}
