package xvelope

const (
	CookieScheme uint8 = 1 << iota
	BearerScheme
	BothScheme = CookieScheme | BearerScheme
)

type Config struct{
	DefaultScheme uint8
	PolicyScheme func(ctx HttpContext) error
	LoginRoute string
	CookieHandler Authenticator
	BearerHandler Authenticator
}

type Auth struct {
	context HttpContext
}

func New( config ...Config) *Auth {
	return &Auth{}
}

func (a *Auth) SetHttpContext( ctx HttpContext) {
	a.context = ctx
}