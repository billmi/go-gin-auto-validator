package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-auto-validator/validator"
)


func main() {

	r := gin.Default()
	indexRouteGroup := r.Group("/index")
	indexRouteGroup.Use()
	{
		indexRouteGroup.POST("/login", Login)
	}
	r.Run(":3000")
}


func Login(c *gin.Context){
	//var phone = c.DefaultPostForm("phone","0")
	//var smsCode = c.DefaultPostForm("smsCode","123456")
	if err, ok := validator.ValiUserLogin(c, "LoginValidator"); ok == false {
		c.JSON(200,gin.H{"error":err})
		return
	}
	return
}
