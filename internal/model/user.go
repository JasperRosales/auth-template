package model

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
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

func (u *User) Build() *User {
	return u
}
