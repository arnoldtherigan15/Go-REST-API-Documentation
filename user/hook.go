package user

import (
	"html"
	"strings"

	"github.com/arnoldtherigan15/Go-REST-API-Documentation/helpers/bcrypt"
)

func (u *User) BeforeSave() error {
	hashedPassword, err := bcrypt.Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))

	return nil
}
