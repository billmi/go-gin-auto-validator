package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

type Validator struct {
	Sence     map[string][]string
	Rules     map[string][]string
	Messages  map[string][]string
	currSence string
}

func (validator *Validator) SetSence(sence string) {
	validator.currSence = sence
}

func (validator *Validator) GetCurrSence() string {
	return validator.currSence
}

func (validator *Validator) GetSenceFields(sence string) []string {
	return validator.Sence[sence]
}

func (validator *Validator) GetRules() map[string][]string {
	tempData := map[string][]string{}
	tempData = validator.Rules
	if validator.currSence != "" {
		tempData = validator.MatchSence(validator.Rules)
	}
	return tempData
}

func (validator *Validator) GetMessages() map[string][]string {
	tempData := map[string][]string{}
	tempData = validator.Messages
	if validator.currSence != "" {
		tempData = validator.MatchSence(validator.Messages)
	}
	return tempData
}

func (validator *Validator) MatchSence(datas map[string][]string) map[string][]string {
	tempData := map[string][]string{}
	senceLists, ok := validator.Sence[validator.GetCurrSence()]
	if ok == false {
		return datas
	}
	for _, row := range senceLists {
		currRow, ok := datas[row]
		if ok == true {
			tempData[row] = currRow
		}
	}
	return tempData
}

func ValiFactory(valiName string) *Validator {
	factory := &Validator{}
	switch valiName {
	case "LoginValidator":
		factory = NewLoginValidator()
		break
	}
	return factory
}

func GetValidator(c *gin.Context, validator *Validator, since string) *govalidator.Validator {
	if since != "" {
		validator.SetSence(since)
	}
	opts := govalidator.Options{
		Request:         c.Request,
		Rules:           validator.GetRules(),
		Messages:        validator.GetMessages(),
		RequiredDefault: true, //所有字段必须通过  true || false
	}
	return govalidator.New(opts)
}


func GetDefaultMsg(values map[string][]string) string {
	temp := ""
	for key, _ := range values {
		row, ok := values[key]
		if ok == true {
			temp = strings.TrimSpace(row[0])
			break
		}
	}
	return temp
}
