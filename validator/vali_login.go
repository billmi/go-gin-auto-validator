package validator

type LoginValidator struct {
	Validator
}

func NewLoginValidator() *Validator {
	loginValidator := &LoginValidator{}
	loginValidator.Sence = map[string][]string{
		"login": []string{"phone", "smsCode"},
	}

	loginValidator.Messages = map[string][]string{
		"phone":  []string{"required:phone", "digits:digits"},
		"smsCode": []string{"required:smsCode"},
	}

	loginValidator.Rules = map[string][]string{
		"phone":  []string{"required", "digits:11"},
		"smsCode": []string{"required"},
	}
	return &loginValidator.Validator
}
