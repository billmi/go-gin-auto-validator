# go-gin-aotu-validator
baseFrom : "github.com/gin-gonic/gin"
baseFrom : "github.com/thedevsaddam/govalidator"

##### 1. get gin
`
go get -u -v github.com/gin-gonic/gin
`

##### 1. get govalidator
`
go get -u -v github.com/thedevsaddam/govalidator
`

#### 实现配置化验证器,验证规则参考 : github.com/thedevsaddam/govalidator

仅仅支持http协议

```
package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-auto-validator/validator"
)

/**
	@author : Bill
	Base : http
	配置解释
	Sence : 场景
	Rules : 规则
	Message : 返回对应消息
 */

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

```