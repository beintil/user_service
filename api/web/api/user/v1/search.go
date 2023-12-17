package v1

import (
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/api/client"
	"github.com/beintil/user_service/api/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestSearchUser represents a request for search a users.
// swagger:request RequestSearchUser
type RequestSearchUser struct {
	// User search form
	client.UserSearchForm
}

// ErrorResponseSearchUser represents a error response for search a users.
// swagger:response ErrorResponseSearchUser
type ErrorResponseSearchUser struct {
	// Error response code
	Code string
	// Error response message
	Message string
}

// SuccessfullyResponseSearchUser represents a successful response for search a users.
// swagger:response SuccessfullyResponseSearchUser
type SuccessfullyResponseSearchUser struct {
	// User Message
	Message string
	// User entity result
	User []client.User
}

// SearchUser
// @Summary Search a users
// @Description Verification of authorization data
// @Tags User
// @Accept json
// @Produce json
// @Param user body RequestSearchUser true "User object to be searched"
// @Success 200 {object} SuccessfullyResponseSearchUser "User search successfully"
// @Failure 400 {object} ErrorResponseSearchUser "Bad request"
// @Failure 500 {object} ErrorResponseSearchUser "Internal server error"
// @Router /api/v1/user/list [post]
func SearchUser(c *gin.Context) {
	form := client.UserSearchForm{}
	c.ShouldBindJSON(&form)
	items, e := core.UsersSearch(db.GetDB(), form)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  e.GetCode(),
			"error": e.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Search user successfully",
		"users":   items,
	})
}
