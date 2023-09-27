package user

import "github.com/gofiber/fiber/v2"

func Apply(r fiber.Router) fiber.Router {
	user := r.Group("/user")
	{

	}
	return user
}
