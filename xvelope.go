package xvelope

import (
	"github.com/elmyrockers/go-xvelope/authenticator/cookie"
	"github.com/elmyrockers/go-xvelope/authenticator/opaque"
	// "github.com/davecgh/go-spew/spew"
)

type Auth struct {
	config Config
	context HttpContext

	authenticator Authenticator
}
func defaultSchemeSelector( HttpContext ) AuthScheme {
	return CookieScheme
}
func New( params ...any) *Auth {
	// Resolve params slice
		var context HttpContext
		if (len(params) > 0){
			context = params[0].(HttpContext)
		}

		var config Config
		if (len(params) > 1){
			config = params[1].(Config)
		}

	// Set authenticator
		var authenticator Authenticator

		// Set scheme
			var scheme AuthScheme

			//CookieAndBearerScheme
				if config.DefaultScheme == 0 || config.DefaultScheme == CookieAndBearerScheme {
					schemeSelector := config.SchemeSelector
					if schemeSelector == nil { schemeSelector = defaultSchemeSelector }
					scheme = schemeSelector( context )
				
			//CookieScheme / BearerScheme
				} else { scheme = config.DefaultScheme }

		// Set handler as authenticator
			if scheme == CookieScheme {
				authenticator = config.CookieHandler
				if authenticator == nil {
					authenticator = cookie.New()
				}
			} else if scheme == BearerScheme {
				authenticator = config.BearerHandler
				if authenticator == nil {
					authenticator = opaque.New()
				}
			}

		// Set httpcontext for authenticator
			if authenticator != nil {
				authenticator.Initialize( context )
			}

	return &Auth{
		config: config,
		context: context,
		authenticator: authenticator,
	}
}

func (a *Auth) SetHttpContext( ctx HttpContext) {
	a.context = ctx
}