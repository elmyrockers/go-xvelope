package fasthttpcontext

import (
	"github.com/valyala/fasthttp"
	"github.com/elmyrockers/go-xvelope"
)

// ------------------------------------------ FastHttpContext
// Compile-time interface checks
var _ HttpContext = (*FastHttpContext)(nil)

type FastHttpContext struct {
	context *fasthttp.RequestCtx
}

func (f *FastHttpContext) SetContext( ctx any ){
	f.context = ctx.(*fasthttp.RequestCtx)
}

func (f *FastHttpContext) Header(key string) string {
	return string(f.context.Request.Header.Peek(key))
}

func (f *FastHttpContext) Cookie(name string) string {
	return string(f.context.Request.Header.Cookie(name))
}

func (f *FastHttpContext) Query(key string) string {
	return string(f.context.QueryArgs().Peek(key))
}

func (f *FastHttpContext) SetStatus(code int) {
	f.context.SetStatusCode(code)
}

func (f *FastHttpContext) SetCookie(cookie *Cookie) {
	if cookie == nil { return }

	// Get cookie instance
		fastCookie := fasthttp.AcquireCookie()
		defer fasthttp.ReleaseCookie(fastCookie)

	// Set cookie
		fastCookie.SetKey(cookie.Name)
		fastCookie.SetValue(cookie.Value)
		fastCookie.SetPath(cookie.Path)
		fastCookie.SetDomain(cookie.Domain)
		fastCookie.SetMaxAge(cookie.MaxAge)
		fastCookie.SetExpire(cookie.Expires)
		fastCookie.SetSecure(cookie.Secure)
		fastCookie.SetHTTPOnly(cookie.HTTPOnly)

		switch cookie.SameSite {
		case SameSiteLax:
			fastCookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
		case SameSiteStrict:
			fastCookie.SetSameSite(fasthttp.CookieSameSiteStrictMode)
		case SameSiteNone:
			fastCookie.SetSameSite(fasthttp.CookieSameSiteNoneMode)
		default:
			fastCookie.SetSameSite(fasthttp.CookieSameSiteDefaultMode)
		}

		f.context.Response.Header.SetCookie(fastCookie)
}

func (f *FastHttpContext) SendJSON(v any) error {
	f.context.SetContentType("application/json")
	return json.NewEncoder(f.context).Encode(v)
}