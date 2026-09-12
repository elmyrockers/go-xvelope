package xvelope

import "github.com/davecgh/go-spew/spew"

type Auth struct {
	config Config
	context HttpContext

	authenticator Authenticator
}

func New( params ...any) *Auth {
	// Resolve params slice
		var context HttpContext
		if (len(params) > 0){
			context = params[0]
		}

		var config Config
		if (len(params) > 1){
			config = params[1]
		}

	// 
	spew.Dump( cfg )

	return &Auth{config: cfg}
}

func (a *Auth) SetHttpContext( ctx HttpContext) {
	a.context = ctx
}