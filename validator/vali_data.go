package validator

import "github.com/gin-gonic/gin"

func valiData(c *gin.Context,validatorName string,sence string) (string,bool) {
	v := GetValidator(c,ValiFactory(validatorName),sence)
	err := v.Validate()
	if len(err) > 0{
		return GetDefaultMsg(err),false
	}
	return "",true
}