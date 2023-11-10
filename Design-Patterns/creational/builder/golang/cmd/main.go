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

	uDirector.BuildRegularUser()

	user := uBuilder.GetUser()

	log.Println(user)
}
