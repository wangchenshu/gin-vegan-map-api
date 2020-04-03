package roles

import (
	"gin-vegan-map-api/server/common"
	e "gin-vegan-map-api/server/enum"
	"gin-vegan-map-api/server/models"
	"gin-vegan-map-api/server/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CheckRoles -
func CheckRoles(roles e.RolesEnum) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.GetQuery("_t")
		if !ok {
			hToken := c.GetHeader("Authorization")
			if len(hToken) < common.BearerLength {
				c.Abort()
				return
			}

			token = strings.TrimSpace(hToken[common.BearerLength:])
			cliams, err := util.ParseToken(token)
			if err != nil {
				log.Println(err)
				c.Abort()
				return
			}

			if !models.CheckUserRoles(cliams.Username, roles) {
				c.JSON(http.StatusForbidden, gin.H{
					"code": http.StatusForbidden,
					"message": e.StatusEnum.String(
						e.NoPermission,
					),
					"data": nil,
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
