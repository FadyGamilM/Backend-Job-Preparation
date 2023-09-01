package user

type Builder struct {
	username string
	password string
}

func (b *Builder) Username(v string) {
	b.username = v
}
func (b *Builder) Password(v string) {
	b.password = v
}

func (b *Builder) Build() *User {
	return &User{
		username: b.username,
		password: b.password,
	}
}
