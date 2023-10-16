package main

import (
	"fmt"

	"github.com/FadyGamilM/designpatterns/user"
)

func main() {
	userBuilder := user.Builder{}
	// once i finished some business logic and i have the final values to create the user object, i need to set the values into the builder and call the Build() method which will return the immutable solidified version of this User object so i can't modify any state of the User states after the creation
	userBuilder.Username("fady")
	userBuilder.Password("fady123456")

	user := userBuilder.Build()
	// now i can only call getters
	fmt.Println(user.GetUsername(), " => ", user.GetPassword())
}
