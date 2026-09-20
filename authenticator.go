package xvelope

import "github.com/elmyrockers/go-xvelope/authenticator"

type Authenticator interface {
	Initialize( context HttpContext )
	Authenticate()
	Challenge()
	Forbid()
	
	SignIn()
	SignOut()
}