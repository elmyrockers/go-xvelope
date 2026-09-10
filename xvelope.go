package xvelope

import "github.com/davecgh/go-spew/spew"

type Auth struct {
	context HttpContext
}

func New( config ...Config) *Auth {
	// Resolve config slice
		var cfg Config
		if len(config) > 0 {
			cfg = config[0]
		}

	// 
	spew.Dump( cfg )

	return &Auth{}
}

func (a *Auth) SetHttpContext( ctx HttpContext) {
	a.context = ctx
}