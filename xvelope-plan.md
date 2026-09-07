## Setup

Import a middleware:
```go
import (
	"github.com/gofiber/fiber/v3"
	fibermw "github.com/elmyrockers/go-xvelope/middleware/fiber"
)
```

Configure:
```go
	app := fiber.New()
	app.Use(fibermw.New(fibermw.Config{}))

	app.Post( "/auth/login",func(c fiber.Ctx) error {
		auth := fibermw.FromContext(c)
	})
```


## Middleware
```go
GuestOnly
AuthRequired
```

