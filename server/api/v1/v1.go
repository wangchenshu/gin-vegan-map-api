package v1

import (
	e "gin-vegan-map-api/server/enum"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrBody -
func ErrBody(code int, message string) gin.H {
	return gin.H{
		"code":    code,
		"message": message,
		"data":    nil,
	}
}

// OkBody -
func OkBody(data interface{}) gin.H {
	return gin.H{
		"code":    http.StatusOK,
		"message": e.StatusEnum.String(e.Success),
		"data":    data,
	}
}
