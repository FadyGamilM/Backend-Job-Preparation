package facade

import (
	"log"
)

type EmailService struct{}

func (es *EmailService) SendEmail() {
	log.Println("Email has been sent successfully .. ")
}

type OtpService struct{}

func (otp *OtpService) Verify() {
	log.Println("OTP verification is done successfully .. ")
}

type AuthService struct{}

func (as *AuthService) Register() {
	log.Println("authentication is done .. ")
}

// the order of execution is that a user recieve an email -> then click the link on the email to recieve an OTP message -> finally back to the register page to click on the register button to be authenticated .. so there is a interdependency between these classes to perform a complex task .. we could simplify it for the client code by creating a facade for it
type AuthFacadeSystem struct {
	emailService *EmailService
	otpService   *OtpService
	authService  *AuthService
}

func NewAuthFacadeSystem(es *EmailService, otps *OtpService, as *AuthService) *AuthFacadeSystem {
	return &AuthFacadeSystem{
		emailService: es,
		otpService:   otps,
		authService:  as,
	}
}

// simplify things with a single api
func (afs *AuthFacadeSystem) Auth() {
	afs.emailService.SendEmail()
	afs.otpService.Verify()
	afs.authService.Register()
}
