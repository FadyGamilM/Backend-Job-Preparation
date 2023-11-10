# Builder Pattern .. ?
- Builder pattern gives your client the highest level of abstraction, because the construction logic separated from the business logic. <br>
➜ the construction logic is set in the `ConcreteBuilder` class <br>
➜ the business logic is simple and its done in the client code via calling the director only. <br>

- Builder pattern gives you the multiple representation, because the director can supports many forms of creating the resource class via the builder. <br>
```go
func (ud *UserDirector) BuildRegularUser(email, username, password string) {
	ud.builder.Email(email)
	ud.builder.Username(username)
	ud.builder.Password(password)
	ud.builder.Role("REGULAR")
}

func (ud *UserDirector) BuildAdminUser() {
	ud.builder.Email("fady@email.com")
	ud.builder.Username("fadygamilm")
	ud.builder.Password("123456789")
	ud.builder.Role("ADMIN")
}
```

- Builder gives us the `object integrity` because we ensure the validity of each object we createed because its a setp-by-step process.

- Builder patterns gives us the `immutability` because we cannot mutate the instance after its created

# When to use and not to use .. ?
- Builder pattern is complex to be implemented and maintained, so unless our object is complex, don't use it.