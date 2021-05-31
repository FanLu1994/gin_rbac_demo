package controller

import (
	"gin_rbac/util"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func GetAuth(c *gin.Context) {
	var user auth
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	// TODO: 数据库验证用户账号密码

	data := make(map[string]interface{})
	code := 200
	token, err := util.GenerateToken(user.Username, user.Password)
	if err != nil {
		code = 400
	} else {
		data["token"] = token
		code = 200
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : 200,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}