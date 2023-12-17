package main

import (
	"github.com/beintil/user_service/api/web"

	_ "github.com/swaggo/http-swagger/example/go-chi/docs"

	"github.com/beintil/user_service/pkg/logger"
	lue "github.com/beintil/user_service/pkg/lu_error"
)

// @title Your API Title User service for the LinkUp application
// @version 1.0
// @host localhost:7081
// @BasePath /user_service/
func main() {
	router := web.GetRoutes()
	err := router.Run("localhost:7081")
	if err != nil {
		logger.Error(lue.New(err, lue.InternalServerError))
	}
}
