package auth

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.SendString("This is login route")
}