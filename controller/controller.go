package controller

import "github.com/gin-gonic/gin"



func GetBook(c *gin.Context){
	c.JSON(200,gin.H{
		"book":[]string{"book1","book2","book3"},
	})
}
