package main



import (
	"github.com/gofiber/fiber/v3"
	fibermw "github.com/elmyrockers/go-xvelope/middleware/fiber"
	"github.com/elmyrockers/go-xvelope"
	// "github.com/davecgh/go-spew/spew"
)




func main() {
	app := fiber.New()
	app.Use(fibermw.New(xvelope.Config{
		DefaultScheme: xvelope.CookieScheme,
	}))

	app.Get( "/",func( c fiber.Ctx ) error {
		auth := fibermw.FromContext(c)
		_ = auth
		// spew.Dump( auth )
		return c.SendString("OK")
	})

	app.Listen(":3000")
}