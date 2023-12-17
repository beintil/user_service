package v1

import (
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/api/client"
	"github.com/beintil/user_service/api/core"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequestCreateUser represents a request for create a user.
// swagger:request Request
type RequestCreateUser struct {
	// User entity
	client.User
}

// ErrorResponseCreateUser represents a error response for create a user.
// swagger:response ErrorResponse
type ErrorResponseCreateUser struct {
	// Error response code
	Code string
	// Error response message
	Message string
}

// SuccessfullyResponseCreateUser represents a successful response for create a user.
// swagger:response SuccessfullyResponseCreateUser
type SuccessfullyResponseCreateUser struct {
	// User Message
	Message string
	// User entity result
	User core.User
}

// CreateUser
// @Summary Create a new user
// @Description Create a new user based on JSON input
// @Tags User
// @Accept json
// @Produce json
// @Param user body RequestCreateUser true "User object to be created"
// @Success 201 {object} SuccessfullyResponseCreateUser "User created successfully"
// @Failure 400 {object} ErrorResponseCreateUser "Bad request"
// @Failure 500 {object} ErrorResponseCreateUser "Internal server error"
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
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
	e := user.Create(db.GetDB())
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  e.GetCode(),
			"error": e.Error(),
		})
		return
	}

	/*s := base.GetAuthService()
	e = s.Login(*user.Email, *user.Password)
	if e != nil {
		logger.Error(e)
	}
	if s.Resp == nil {
		logger.Error(lueNew(luePortErrorSystem, "System error connect authorization service").HTTP(http.StatusInternalServerError))
	} else {
		cookies := s.Resp.Cookies()
		if len(cookies) <= 0 {
			logger.Error(lueNew(luePortErrorSystem, fmt.Sprint("System error set auth cookie: ", err.Error())).HTTP(http.StatusNoContent))
		} else {
			for i := range cookies {
				http.SetCookie(w, cookies[i])
			}
		}
	}*/

	c.JSON(http.StatusCreated, gin.H{
		"message": "Create user successfully",
		"user":    user,
	})
	//gorest.Send(w, gorest.NewOkJsonResponse("Create User successfully", user, nil))
}
