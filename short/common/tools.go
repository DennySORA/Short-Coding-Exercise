package common

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PKGCloseFunc(function func()) func() error {
	return func() error {
		function()
		return nil
	}
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func CheckExpires(timeExpires int64) (bool, error) {
	return time.Now().Unix() > timeExpires, nil
}

func GetShortError() (int, gin.H) {
	statusCode := http.StatusNotFound
	description := gin.H{
		"description": "Parameter Error.",
	}
	return statusCode, description
}
