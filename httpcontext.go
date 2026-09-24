package xvelope

import "time"

type SameSite int

const (
	SameSiteDefault SameSite = iota
	SameSiteLax
	SameSiteStrict
	SameSiteNone
)

type Cookie struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	MaxAge   int
	Expires  time.Time
	Secure   bool
	HTTPOnly bool
	SameSite SameSite
}

type HttpContext interface {
	// --- Reading the incoming credential ---
	Header(key string) string           // "Authorization: Bearer <token>"
	Cookie(name string) string 			// "auth_payload" cookie
	Query(key string) string            // "?useCookies=true" - detect cookie or token based

	// --- Writing the response ---
	SetStatus(code int)
	SetCookie(cookie *Cookie)
	SendJSON(v any) error

	SetContext( ctx any )
}