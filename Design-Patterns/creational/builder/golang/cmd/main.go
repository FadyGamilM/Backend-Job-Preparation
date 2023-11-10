package main

import (
	user "builder_pattern/User"
	"log"
)

func main() {
	// instantiate an instance of the concrete builder
	uBuilder := user.NewConcreteUserBuilder()

	// instantiate an instance of the director
	uDirector := user.NewUserDirector(uBuilder)

	uDirector.BuildRegularUser("fady@mail.con", "fgm", "123456789")

	user := uBuilder.GetUser()

	log.Println(user)
}
