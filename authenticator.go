package xvelope

type Authenticator interface {
	Authenticate()
	Challenge()
	Forbid()
	
	SignIn()
	SignOut()
}