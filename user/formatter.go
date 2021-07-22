package user

type Formatter struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func FormatUser(user *User) Formatter {
	formatted := Formatter{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  *user.LastName,
	}
	return formatted
}
