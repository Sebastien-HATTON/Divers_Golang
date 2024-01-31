package main

import (
	"github.com/gofiber/fiber/v2"
)

const boat = `
    ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/
`


func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(boat + "\n   Hello World !")
	})

	app.Listen(":8081")
}
