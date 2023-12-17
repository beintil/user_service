package web

import (
	"fmt"
	"github.com/beintil/user_service/api/client"
	userV1 "github.com/beintil/user_service/api/web/api/user/v1"
	"github.com/beintil/user_service/pkg"
	"github.com/beintil/user_service/pkg/logger"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"net/http"
	"os"
	"path/filepath"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())

	// user service path
	mainPath := router.Group("user_service")
	// api service
	api := mainPath.Group("api")
	// path v1
	apiV1 := api.Group("v1", httpLogger())

	// User sub router
	userV1.Init(apiV1)

	return router
}

func httpLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logFile, _ := os.OpenFile(createLogFile(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		var log = make(map[string]interface{})
		log["PATH"] = ctx.FullPath()
		log["CODE"] = ctx.Writer.Status()
		log["CODE_TEXT"] = http.StatusText(ctx.Writer.Status())
		log["ERROR"] = ctx.Errors.String()
		if ctx.Writer.Status() == http.StatusInternalServerError {
			logFile.WriteString(pkg.WithJSON(log))
			fmt.Println(client.ErrorColor.Sprintln(pkg.WithJSON(log)))
		} else {
			logger.Info(pkg.WithJSON(log))
		}
	}
}

func createLogFile() string {
	logFilePath := "./var/logs/http.log"
	if err := os.MkdirAll(filepath.Dir(logFilePath), 0755); err != nil {
		logger.Error(lue.New(client.ErrorColor.Sprint("Unable to create log directory: ", err.Error()), lue.InternalServerError))
		return ""
	}

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Error(lue.New(client.ErrorColor.Sprint("Unable to open log file: ", err.Error()), lue.InternalServerError))
		return ""
	}
	return logFile.Name()
}
