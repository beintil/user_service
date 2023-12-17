package v1

import (
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/api/client"
	"github.com/beintil/user_service/api/core"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestVerify represents a request for check auth data the user.
// swagger:request RequestVerify
type RequestVerify struct {
	// Auth data for the user
	client.AuthData
}

// ErrorResponseVerify represents a error response for check auth data the user.
// swagger:response ErrorResponseVerify
type ErrorResponseVerify struct {
	// Error response code
	Code string
	// Error response message
	Message string
}

// SuccessfullyResponseVerify represents a successful response for check auth data the user.
// swagger:response SuccessfullyResponseVerify
type SuccessfullyResponseVerify struct {
	// User Message
	Message string
	// User entity result
	User core.User
}

// Verify
// @Summary Check auth data the user
// @Description Check auth data the user based on JSON input
// @Tags User
// @Accept json
// @Produce json
// @Param user body RequestVerify true "User object to be checked"
// @Success 200 {object} SuccessfullyResponseVerify "User load successfully"
// @Failure 400 {object} ErrorResponseVerify "Bad request"
// @Failure 500 {object} ErrorResponseVerify "Internal server error"
// @Router /api/v1/user/verify [post]
func Verify(c *gin.Context) {
	user := core.NewUser()
	err := c.ShouldBindJSON(&user.User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  lue.CodeInText[lue.BadRequest],
			"error": err.Error(),
		})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  lue.CodeInText[lue.BadRequest],
			"error": "Entity in the request is empty",
		})
		return
	}
	isCorrect := user.Verify(db.GetDB())
	if !isCorrect {
		c.JSON(http.StatusProxyAuthRequired, gin.H{
			"code":  lue.CodeInText[lue.Unauthorized],
			"error": "Invalid auth data",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Load user successfully",
		"user":    user,
	})
}
