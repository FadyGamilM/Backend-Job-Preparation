package user

// abstract interface that describes how to build a resource
type IUserBuilder interface {
	Username(v string)
	Email(v string)
	Password(v string)
	Role(v string)
}

type User struct {
	username string
	email    string
	password string
	role     string
}

// the concrete User Builder implements the IUserBuilder
type concreteUserBuilder struct {
	user *User // i defined it as pointer because it might be nil
}

func NewConcreteUserBuilder() *concreteUserBuilder {
	return &concreteUserBuilder{
		user: &User{}, // initialize the memory with a stored space with default values, so later we can call .Email() and set properties on this object reference
	}
}

func (cub *concreteUserBuilder) Username(v string) {
	cub.user.username = v
}

func (cub *concreteUserBuilder) Email(v string) {
	cub.user.email = v
}

func (cub *concreteUserBuilder) Password(v string) {
	cub.user.password = v
}

func (cub *concreteUserBuilder) Role(v string) {
	cub.user.role = v
}

// method to get the built resource
func (cub *concreteUserBuilder) GetUser() *User {
	return cub.user
}

// Director: The director is responsible for overseeing the construction of the product. It works with the Builder interface to build the product step-by-step.
type UserDirector struct {
	builder IUserBuilder
}

func NewUserDirector(b IUserBuilder) *UserDirector {
	return &UserDirector{
		builder: b,
	}
}

func (ud *UserDirector) BuildRegularUser(email, username, password string) {
	ud.builder.Email(email)
	ud.builder.Username(username)
	ud.builder.Password(password)
	ud.builder.Role("REGULAR")
}

func (ud *UserDirector) BuildAdminUser(email, username, password string) {
	ud.builder.Email(email)
	ud.builder.Username(username)
	ud.builder.Password(password)
	ud.builder.Role("ADMIN")
}
