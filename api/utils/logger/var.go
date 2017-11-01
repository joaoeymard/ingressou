package logger

import (
	"strings"
	"sync"
)

var (
	mutex sync.Mutex
)

// Status Retorna o status da requesição
func Status(response string) string {

	if strings.Contains(response, "&") {
		response = strings.Replace(response, "&", "", 1)
		columns := strings.Split(response, " ")
		for _, column := range columns {
			if strings.Contains(strings.ToLower(column), "status") {
				return strings.Split(column, ":")[1]
			}
		}
	}
	return ""
}
