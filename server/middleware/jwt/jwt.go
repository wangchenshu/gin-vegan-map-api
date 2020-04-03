package jwt

import (
	e "gin-vegan-map-api/server/enum"
	"net/http"
	"strings"
	"time"

	"gin-vegan-map-api/server/util"

	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT(c *gin.Context) {
	var (
		data    interface{}
		retCode int
		retMsg  string
	)

	retCode = e.Success
	token := c.DefaultQuery("token", "")
	if token == "" {
		token = c.Request.Header.Get("Authorization")
		if s := strings.Split(token, " "); len(s) == 2 {
			token = s[1]
		}
	}

	if token == "" {
		retCode = http.StatusBadRequest
		retMsg = e.StatusEnum.String(e.InvalidParams)
	} else {
		claims, err := util.ParseToken(token)
		if err != nil {
			retCode = e.CheckTokenFail
			retMsg = e.AuthEnum.String(e.CheckTokenFail)
		} else if time.Now().Unix() > claims.ExpiresAt {
			retCode = e.TokenExpired
			retMsg = e.AuthEnum.String(e.TokenExpired)
		}
	}

	if retCode != e.Success {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": retCode,
			"msg":  retMsg,
			"data": data,
		})

		c.Abort()
		return
	}

	c.Next()
}
