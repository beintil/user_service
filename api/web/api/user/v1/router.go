package v1

import (
	"github.com/gin-gonic/gin"
)

// User sub route
func Init(apiRoute *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	user := apiRoute.Group("user")
	// Create a new user
	user.POST("", CreateUser)
	// Load the user by id
	user.GET("/:id", LoadUser)
	// Search a users
	user.POST("/list", SearchUser)
	// Verification of authorization data
	user.POST("/verify", Verify)

	user.Use(middleware...)
}
