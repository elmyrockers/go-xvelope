package xvelope

type AuthScheme uint8

const (
	CookieScheme AuthScheme = 1 << iota
	BearerScheme
	CookieAndBearerScheme = CookieScheme | BearerScheme
)

type Config struct{
	DefaultScheme AuthScheme
	SchemeSelector func(ctx HttpContext) AuthScheme
	LoginRoute string
	CookieHandler Authenticator
	BearerHandler Authenticator
}