package middleware

import (
	"gin_rbac/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc{
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Request.Header.Get("Authorization")
		if token==""{
			code = 400
		}else{
			claims,err := util.ParseToken(token)
			if err!=nil{
				code = 404
			}else if time.Now().Unix()>claims.ExpiresAt{
				code = 501
			}
		}

		if code != 200{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code" : code,
				"msg" : "需要登录",
				"data" : data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}