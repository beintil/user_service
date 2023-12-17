package v1

import (
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/api/core"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RequestLoadUser represents a request for load the user.
// swagger:request RequestUser
type RequestLoadUser struct {
	// Identification of the user
	Id int
}

// ErrorResponseLoadUser represents a error response for load the user.
// swagger:response ErrorResponse
type ErrorResponseLoadUser struct {
	// Error response code
	Code string
	// Error response message
	Message string
}

// SuccessfullyResponseLoadUser represents a successful response for load the user.
// swagger:response SuccessfullyResponseLoadUser
type SuccessfullyResponseLoadUser struct {
	// User Message
	Message string
	// User entity result
	User core.User
}

// LoadUser
// @Summary Load the user
// @Description Load the user based on JSON input
// @Tags User
// @Accept json
// @Produce json
// @Param user body RequestLoadUser true "User object to be loaded"
// @Success 200 {object} SuccessfullyResponseLoadUser "User load successfully"
// @Failure 400 {object} ErrorResponseLoadUser "Bad request"
// @Failure 500 {object} ErrorResponseLoadUser "Internal server error"
// @Router /api/v1/user/{id} [get]
func LoadUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  lue.CodeInText[lue.BadRequest],
			"error": err.Error(),
		})
		return
	}
	user := core.NewUser().SetPrimary(id)
	e := user.Load(db.GetDB())
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  e.GetHTTP(),
			"error": e.Error(),
		})
		return
	}
	user.ClearSensitiveData()
	c.JSON(http.StatusOK, gin.H{
		"message": "Load user successfully",
		"user":    user,
	})
}
