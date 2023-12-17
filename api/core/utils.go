package core

import "fmt"

func SetAlias(alias string, columns ...string) []string {
	for i, column := range columns {
		columns[i] = fmt.Sprint(alias, ".", column)
	}
	return columns
}
