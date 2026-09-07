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
	app.Use(fibermw.New(fibermw.Config{

	}))
```


## Middleware
```go
GuestOnly
AuthRequired
```

