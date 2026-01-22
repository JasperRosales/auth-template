package model

import "fmt"

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Token	string `json:"token,omitempty"`
}

func New() *User {
	return &User{}
}

func (u *User) SetId(id int64) *User {
	u.ID = id
	return u
}

func (u *User) SetEmail(email string) *User {
	u.Email = email
	return u
}

func (u *User) SetPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) SetToken(token string) *User {
	u.Token = token
	return u
}

func (u *User) Build() *User {
	return u
}


func (u *User) PrettyPrint() string {
	return "User {\n" +
		"  ID: " + fmt.Sprintf("%d", u.ID) + "\n" +
		"  Email: " + u.Email + "\n" +
		"  Password: " + u.Password + "\n" +
		"  Token: " + u.Token + "\n" +
		"}"
}
