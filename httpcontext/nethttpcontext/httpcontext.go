package nethttpcontext

import (
	"net/http"
	"encoding/json"

	"github.com/elmyrockers/go-xvelope"
)

// ------------------------------------------ NetHttpContext
// Compile-time interface checks
var	_ HttpContext = (*NetHttpContext)(nil)

type NetHttpContextPair struct {
	Request *http.Request
	Response http.ResponseWriter
}

type NetHttpContext struct {
	context NetHttpContextPair
}

func (n *NetHttpContext) SetContext( ctx any ){
	n.context = ctx.(NetHttpContextPair)
}

func (n *NetHttpContext) Header(key string) string {
	return n.context.Request.Header.Get(key)
}

func (n *NetHttpContext) Cookie(name string) string {
	c, err := n.context.Request.Cookie(name)
	if err != nil { return "" }
	return c.Value
}

func (n *NetHttpContext) Query(key string) string {
	return n.context.Request.URL.Query().Get(key)
}

func (n *NetHttpContext) SetStatus(code int) {
	n.context.Response.WriteHeader(code)
}

func (n *NetHttpContext) SetCookie(cookie *Cookie) {
	if cookie == nil { return }
	
	var sameSite http.SameSite
	switch cookie.SameSite {
	case SameSiteStrict:
		sameSite = http.SameSiteStrictMode
	case SameSiteNone:
		sameSite = http.SameSiteNoneMode
	case SameSiteLax:
		sameSite = http.SameSiteLaxMode
	default:
		sameSite = http.SameSiteDefaultMode
	}

	http.SetCookie(n.context.Response, &http.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Path:     cookie.Path,
		Domain:   cookie.Domain,
		MaxAge:   cookie.MaxAge,
		Expires:  cookie.Expires,
		Secure:   cookie.Secure,
		HttpOnly: cookie.HTTPOnly,
		SameSite: sameSite,
	})
}

func (n *NetHttpContext) SendJSON(v any) error {
	n.context.Response.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(n.context.Response).Encode(v)
}