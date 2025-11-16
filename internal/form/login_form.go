package form

import "net/mail"

type LoginForm struct {
	Email string
}

func (f *LoginForm) Validate() map[string]string {
	errors := make(map[string]string)

	if f.Email == "" {
		errors["Email"] = "Email is required"
	} else if _, err := mail.ParseAddress(f.Email); err != nil {
		errors["Email"] = "Email must be a valid email address"
	}

	return errors
}
