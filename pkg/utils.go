package pkg

import (
	"fmt"
	"strings"
)

func WithJSON(values map[string]interface{}) string {
	var builder strings.Builder
	builder.WriteString("\n\n{\n")

	for k, v := range values {
		builder.WriteString("  ")
		builder.WriteString(k)
		builder.WriteString(":\n")
		builder.WriteString("   ")
		builder.WriteString(fmt.Sprint(v))
		builder.WriteString("\n")
	}
	builder.WriteString("}")
	return builder.String()
}
