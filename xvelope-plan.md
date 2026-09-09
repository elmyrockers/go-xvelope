## Setup

Import a middleware:
```go
import (
	"github.com/gofiber/fiber/v3"
	"github.com/elmyrockers/go-xvelope"
	fibermw "github.com/elmyrockers/go-xvelope/middleware/fiber"
)
```

Configure:
```go
	app := fiber.New()
	app.Use(fibermw.New(xvelope.Config{
		DefaultScheme: xvelope.CookieAndBearerScheme,
		SchemeSelector: nil,
		LoginRoute: nil,
		CookieHandler: nil,
		BearerHandler: nil,
	}))

	app.Post( "/auth/login",func(c fiber.Ctx) error {
		auth := fibermw.FromContext(c)
	})
```

## Middleware
```go
GuestOnly
AuthRequired (RequireAuthorization)
```