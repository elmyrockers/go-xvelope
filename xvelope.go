package xvelope

type Auth struct {
	context HttpContext
}

func New( config ...Config) *Auth {
	return &Auth{}
}

func (a *Auth) SetHttpContext( ctx HttpContext) {
	a.context = ctx
}