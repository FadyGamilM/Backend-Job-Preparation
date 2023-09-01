package user

type User struct {
	username string
	password string
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetPassword() string {
	return u.password
}
