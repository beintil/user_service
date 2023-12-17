package logger

import (
	"fmt"
	"github.com/beintil/user_service/api/client"
	"github.com/beintil/user_service/pkg/lu_error"
	"net/http"
	"os"
	"runtime"
)

type Logger struct {
	skip int
}

func NewLogger(skip int) *Logger {
	return &Logger{skip: skip}
}

func Error(err lue.IError) {
	NewLogger(2).Error(err)
}

func Info(info ...string) {
	NewLogger(2).Info(info...)
}

func (l *Logger) Error(err lue.IError) {
	if l.skip == 0 {
		l.skip = 1
	}
	_, file, line, _ := runtime.Caller(l.skip)
	e := fmt.Sprintf("\n{\n  File: \n   %s\n  Line: \n   %d\n  Error:\n   %v\n  Code:\n   %s\n  HTTP:\n   %v\n}", file, line, err.Error(), err.GetCode(), http.StatusText(err.GetHTTP()))
	fmt.Println(client.ErrorColor.Sprint(e))
	l.saveFile(e)
}

func (l *Logger) Info(info ...string) {
	fmt.Println(client.InformationColorYellow.Sprintln("Information: ", info))
}

func (l *Logger) saveFile(info string) {
	logFile, err := os.OpenFile("./var/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(client.ErrorColor.Sprint("Unable to open log file: ", err.Error()))
	}
	defer logFile.Close()

	logFile.WriteString(info)
}
