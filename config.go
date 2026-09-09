package xvelope

const (
	CookieScheme uint8 = 1 << iota
	BearerScheme
	CookieAndBearerScheme = CookieScheme | BearerScheme
)

type Config struct{
	DefaultScheme uint8
	PolicyScheme func(ctx HttpContext) error
	LoginRoute string
	CookieHandler Authenticator
	BearerHandler Authenticator
}