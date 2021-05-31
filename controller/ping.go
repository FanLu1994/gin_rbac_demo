package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/casbin/casbin/v2"
	"log"
	"os"
	"path"
)


func Ping(c *gin.Context) {
	str, _ := os.Getwd()
	confPath := path.Join(str,"/casbin/model.,conf")
	csvPath := path.Join(str,"/casbin/policy.csv")

	e, err := casbin.NewEnforcer(confPath,csvPath)
	//fmt.Print(e.GetPolicy())
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	sub := "bob" // the user that wants to access a resource.
	obj := "data2" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.
	//  注意点：这里的数据一把都是从数据库拿到的 ，数据库里记录了模型文件对应的策略信息。Casbin 根据策略字段判断是否有权限

	var msg string
	result,_ := e.Enforce(sub,obj,act)

	if result == true {
		fmt.Println("进入了")
		msg = "进入了"
	} else {
		fmt.Println("拒绝了")
		msg = "无权访问"
	}

	c.JSON(200, gin.H{
		"message": msg,
	})
}
