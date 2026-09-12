package xvelope

type Authenticator interface {
	Initialize( context HttpContext )
	Authenticate()
	Challenge()
	Forbid()
	
	SignIn()
	SignOut()
}