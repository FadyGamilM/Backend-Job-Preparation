package main

import "facade/facade"

func main() {
	emailService := &facade.EmailService{}
	otpService := &facade.OtpService{}
	registerationService := &facade.AuthService{}
	authFacade := facade.NewAuthFacadeSystem(emailService, otpService, registerationService)
	authFacade.Auth()
}

/*
	[OUTPUT]
	➜ facade git:(creational-design-patterns) go run .\cmd\main.go
		2023/11/10 10:05:05 Email has been sent successfully ..
		2023/11/10 10:05:05 OTP verification is done successfully ..
		2023/11/10 10:05:05 authentication is done ..
*/
