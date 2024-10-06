package login

import "rachitmishra.com/pastebin/internal/validator"

type LoginForm struct {
	username string
	password string
	validator.Validator
}

func (p *LoginForm) Hash() string {
	return ""
}

type LoginVM struct {
	form *LoginForm
}

func NewLoginVM(pasteForm *LoginForm) LoginVM {
	return LoginVM{
		form: pasteForm,
	}
}
