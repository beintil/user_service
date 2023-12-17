package main

import (
	"fmt"
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/pkg/logger"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"os"
	"strings"
)

var path = "./api/client/business_constants.go"

func main() {
	os.Remove(path)
	file, err := os.Create(path)
	if err != nil {
		logger.Error(lue.New("CREATE FILE ERROR "+err.Error(), lue.ErrorCreate))
		return
	}
	defer file.Close()

	rows, err := db.GetDB().Query("SELECT type, label, code, info FROM business_constants ORDER BY code")
	if err != nil {
		os.Remove(path)
		logger.Error(lue.New("ERROR GET FROM DB business_constants "+err.Error(), lue.ErrorCreate))
		return
	}
	defer rows.Close()

	var builder strings.Builder
	builder.WriteString("package client\n\n")
	builder.WriteString("const (\n")
	for rows.Next() {
		var code int
		var label string
		var typeConst string
		var info string
		if err = rows.Scan(&typeConst, &label, &code, &info); err != nil {
			os.Remove(path)
			logger.Error(lue.New("ERROR SCAN "+err.Error(), lue.InternalServerError))
			return
		}
		builder.WriteString(fmt.Sprintf("\t%s_%s = %d // %s\n", strings.ToUpper(typeConst), strings.ToUpper(label), code, info))
	}
	builder.WriteString(")")

	_, err = fmt.Fprintln(file, builder.String())
	if err != nil {
		os.Remove(path)
		logger.Error(lue.New(fmt.Sprintln("ERROR WRITE FOR WILE ", file.Name(), err.Error()), lue.ErrorWriter))
		return
	}

	logger.Info("GENERATE CONSTANTS SUCCESSFULLY")
}
